var usersApp = angular.module('usersApp', []);
usersApp.controller('UsersCtrl', function ($scope, $http){
	$http.get('/v1/users').success(function(data) {
		$scope.users = data;
	});
});
