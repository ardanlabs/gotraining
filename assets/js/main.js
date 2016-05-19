(function() {
  // Initialize highlight.js
  hljs.initHighlightingOnLoad();

  // Add bootstrap classes on tables
  $("table").addClass("table table-bordered");

  $(".run").click(function() {

    var file = $(this).data("file");
    var output = $(".output");

    // Hide the results if they were ran previously
    output.text("").addClass("hidden").removeClass("alert alert-danger");

    var handle = function(text, success) {

      // Highlight errors
      if (!success) {
        output.addClass("alert alert-danger");
      }

      // Add output text to the dom and show the result
      output.text(text).removeClass("hidden");

      // Scroll the window down to show the output since it will probably have
      // been revealed below the visible space.
      $('html, body').animate({scrollTop: output.offset().top}, 1000);
    };

    // Run the file then handle the response. The function signatures for
    // `done` and `fail` are different so we wrap the main behavior in `handle`
    $.post("/run/" + file)
      .done(function(data) {
        handle(data, true);
      })
      .fail(function(resp) {
        handle(resp.responseText, false);
      });
  });
})()
