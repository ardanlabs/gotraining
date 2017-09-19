# Changelog

## 1.5.1 / 1.5.2

### Bug Fixes

* A pipeline can get stuck after repeated worker failures.  (#2064)
* `pachctl port-forward` can leave a orphaned process after it exits.  (#2098)
* `alpine`-based pipelines fail to load input data.  (#2118)
* Logs are written to the object store even when stats is not enabled, slowing down the pipeline unnecessarily.  (#2119)

### Features / Improvements

* Pipelines now support the “stats” feature.  See the [docs](http://pachyderm.readthedocs.io/en/latest/reference/pipeline_spec.html#enable-stats-optional) for details.  (#1998)
* Pipeline cache size is now configurable.  See the [docs](http://pachyderm.readthedocs.io/en/latest/reference/pipeline_spec.html#cache-size-optional) for details.  (#2033)
* `pachctl update-pipeline` now **only** process new input data with the new code; the old input data is not re-processed.  If it’s desired that all data are re-processed, use the `--reprocess` flag.  See the [docs](http://pachyderm.readthedocs.io/en/latest/fundamentals/updating_pipelines.html) for details.  (#2034)
* Pipeline workers now support “pipelining”, meaning that they start downloading the next datums while processing the current datum, thereby improving overall throughput.  (#2057)
* The `scaleDownThreshold` feature has been improved such that when a pipeline is scaled down, the remaining worker only takes up minimal system resources.  (#2091)

## 1.5.0

### Bug Fixes

* Downstream repos' provenance is not updated properly when `update-pipeline` changes the inputs for a pipeline. (#1958)
* `pachctl version` blocks when pachctl doesn't have Internet connectivity. (#1971)
* `incremental` misbehaves when files are deeply nested. (#1974)
* An `incremental` pipeline blocks if there's provenance among its inputs. (#2002)
* PPS fails to create subsequent pipelines if any pipeline failed to be created. (#2004)
* Pipelines sometimes reprocess datums that have already been processed. (#2008)
* Putting files into open commits fails silently. (#2014)
* Pipelines with inputs that use different branch names fail to create jobs. (#2015)
* `get-logs` returns incomplete logs.  (#2019)

### Features

* You can now use `get-file` and `list-file` on open commits. (#1943)

## 1.4.8

### Bug Fixes

- Fixes bugs that caused us to swamp etcd with traffic.
- Fixes a bug that could cause corruption to in pipeline output.

### Features

- Readds incremental processing mode
- Adds `DiffFile` which is similar in function to `git diff`
- Adds the ability to use cloudfront as a caching layer for additional scalability on aws.
- `DeletePipeline` now allows you to delete the output repos as well.
- `DeletePipeline` and `DeleteRepo` now support a `--all` flag

### Removed Features

- Removes one-off jobs, they were a rarely used feature and the same behavior can be replicated with pipelines

## 1.4.7

### Bug fixes

* [Copy elision](http://pachyderm.readthedocs.io/en/latest/reference/best_practices.html#shuffling-files) does not work for directories. (#1803)
* Deleting a file in a closed commit fails silently. (#1804)
* Pachyderm has trouble processing large files. (#1819)
* etcd uses an unexpectedly large amount of space. (#1824)
* `pachctl mount` prints lots of benevolent FUSE errors. (#1840)

### New features

* `create-repo` and `create-pipeline` now accept the `--description` flag, which creates the repo/pipeline with a "description" field.  You can then see the description via `inspect-repo/inspect-pipeline`. (#1805)
* Pachyderm now supports garbage collection, i.e. removing data that's no longer referenced anywhere.  See the [docs](http://pachyderm.readthedocs.io/en/latest/reference/best_practices.html#garbage-collection) for details. (#1826)
* Pachyderm now has GPU support!  See the [docs](http://pachyderm.readthedocs.io/en/latest/cookbook/tensorflow_gpu.html) for details. (#1835)
* Most commands in `pachctl` now support the `--raw` flag, which prints the raw JSON data as opposed to pretty-printing.  For instance, `pachctl inspect-pipeline --raw` would print something akin to a pipeline spec. (#1839)
* `pachctl` now supports `delete-commit`, which allows for deleting a commit that's not been finished.  This is useful when you have added the wrong data in a commit and you want to start over.
* The web UI has added a file viewer, which allows for viewing PFS file content in the browser.

## 1.4.6

### Bug fixes

* `get-logs` returns errors along the lines of `Invalid character…`. (#1741)
* etcd is not properly namespaced. (#1751)
* A job might get stuck if it uses `cp -r` with lazy files. (#1757)
* Pachyderm can use a huge amount of memory, especially when it processes a large number of files. (#1762)
* etcd returns `database space exceeded` errors after the cluster has been running for a while. (#1771)
* Jobs crashing might eventually lead to disk space being exhausted. (#1772)
* `port-forward` uses wrong port for UI websocket requests to remote clusters (#1754)
* Pipelines can end up with no running workers when the cluster is under heavy load. (#1788)
* API calls can start returning `context deadline exceeded` when the cluster is under heavy load. (#1796)

### New features / improvements

* Union input: a pipeline can now take the union of inputs, in addition to the cross-product of them.  Note that the old `inputs` field in the pipeline spec has been deprecated in favor of the new `input` field.  See the [pipeline spec](http://pachyderm.readthedocs.io/en/latest/reference/pipeline_spec.html#input-required) for details. (#1665)
* Copy elision: a pipeline that shuffles files can now be made more efficient by simply outputting symlinks to input files.  See the [docs on shuffling files](http://pachyderm.readthedocs.io/en/latest/reference/best_practices.html#shuffling-files) for details. (#1791)
* `pachctl glob-file`: ever wonder if your glob pattern actually works?  Wonder no more.  You can now use `pachctl glob-file` to see the files that match a given glob pattern. (#1795)
* Workers no longer send/receive data through pachd.  As a result, pachd is a lot more responsive and stable even when there are many ongoing jobs.  (#1742)

## 1.4.5

### Bug fixes

* Fix a bug where pachd may crash after creating/updating a pipeline that has many input commits. (#1678)
* Rules for determining when input data is re-processed are made more intuitive.  Before, if you update a pipeline without updating the `transform`, the input data is not re-processed.  Now, different pipelines or different versions of pipelines always re-process data, even if they have the same `transform`. (#1685)
* Fix several issues with jobs getting stuck. (#1717)
* Fix several issues with lazy pipelines getting stuck. (#1721)
* Fix an issue with Minio deployment that results in job crash loop. (#1723)
* Fix an issue where a job can crash if it outputs a large number of files. (#1724)
* Fix an issue that causes intermittent gRPC errors. (#1727)

### New features

* Pachyderm now ships with a web UI!  To deploy a new Pachyderm cluster with the UI, use `pachctl deploy <arguments> --dashboard`.  To deploy the UI onto an existing cluster, use `pachctl deploy <arguments> --dashboard-only`.  To access the UI, simply `pachctl port-forward`, then go to `localhost:38080`.  Note that the web UI is currently in alpha; expect bugs and significant changes.   
* You can now specify the amount of resources (i.e. CPU & memory) used by Pachyderm and etcd.  See `pachctl deploy --help` for details. (#1676)
* You can now specify the amount of resources (i.e. CPU & memory) used by your pipelines.  See the [pipeline spec](http://pachyderm.readthedocs.io/en/latest/reference/pipeline_spec.html#resource-spec-optional) for details. (#1683)

## 1.4.4

### Bug fixes

* A job can fail to restart when encountering an internal error.
* A deployment with multiple pachd nodes can get stalled jobs.
* `delete-pipeline` is supposed to have the `--delete-jobs` flag but doesn't.
* `delete-pipeline` can fail if there are many jobs in the pipeline.
* `update-pipeline` can fail if the original pipeline has not outputted any commits.
* pachd can crash if etcd is flaky.
* pachd memory can be easily exhausted on GCE deployments.
* If a pipeline is created with multiple input commits already present, all jobs spawn and run in parallel.  After the fix, jobs always run serially.

### Features

* Pachyderm now supports auto-scaling: a pipeline's worker pods can be terminated automatically when the pipeline has been idle for a configurable amount of time.  See the `scaleDownThreshold` field of the [pipeline spec](http://pachyderm.readthedocs.io/en/latest/reference/pipeline_spec.html#scale-down-threshold-optional) for details.
* The processing of a datum can be restarted manually via `restart-datum`.
* Workers' statuses are now exposed through `inspect-job`.
* A job can be stopped manually via `stop-job`.

## 1.4.3

### Bug fixes

* Pipelines with multiple inputs process only a subset of data.
* Workers may fall into a crash loop under certain circumstances. (#1606)

### New features

* `list-job` and `inspect-job` now display a job's progress, i.e. they display the number of datums processed thus far, and the total number of datums.
* `delete-pipeline` now accepts an option (`--delete-jobs`) that deletes all jobs in the pipeline. (#1540)
* Azure deployments now support dynamic provisioning of volumes.

## 1.4.2

### Bug fixes

* Certain network failures may cause a job to be stuck in the `running` state forever.
* A job might get triggered even if one of its inputs is empty.
* Listing or getting files from an empty output commit results in `node "" not found` error.
* Jobs are not labeled as `failure` even when the user code has failed.
* Running jobs do not resume when pachd restarts.
* `put-file --recursive` can fail when there are a large number of files.
* minio-based deployments are broken.

### Features

* `pachctl list-job` and `pachctl inspect-job` now display the number of times each job has restarted.
* `pachctl list-job` now displays the pipeline of a job even if the job hasn't completed.

## 1.4.1

### Bug fixes

* Getting files from GCE results in errors.
* A pipeline that has multiple inputs might place data into the wrong `/pfs` directories.
* `pachctl put-file --split` errors when splitting to a large number of files.
* Pipeline names do not allow underscores.
* `egress` does not work with a pipeline that outputs a large number of files. 
* Deleting nonexistent files returns errors.
* A job might try to process datums even if the job has been terminated.
* A job doesn't exit after it has encountered a failure.
* Azure backend returns an error if it writes to an object that already exists.

### New features

* `pachctl get-file` now supports the `--recursive` flag, which can be used to download directories.
* `pachctl get-logs` now outputs unstructured logs by default.  To see structured/annotated logs, use the `--raw` flag.

## 1.4.0

Features/improvements:

- Correct processing of modifications and deletions.  In prior versions, Pachyderm pipelines can only process data additions; data that are removed or modified are effectively ignored.  In 1.4, when certain input data are removed (or modified), downstream pipelines know to remove (or modify) the output that were produced as a result of processing the said input data.

As a consequence of this change, a user can now fix a pipeline that has processed erroneous data by simply making a new commit that fixes the said erroneous data, as opposed to having to create a new pipeline.

- Vastly improved performance for metadata operations (e.g. list-file, inspect-file).  In prior versions, metadata operations on commits that are N levels deep are O(N) in runtime.  In 1.4, metadata operations are always O(1), regardless of the depth of the commit. 

- A new way to specify how input data is partitioned.  Instead of using two flags `partition` and `incrementality`, we now use a single `glob` pattern.  See the [glob doc](http://pachyderm.readthedocs.io/en/stable/reference/pipeline_spec.html#input-glob-pattern) for details.

- Flexible branch management.  In prior versions, branches are fixed, in that a commit always stays on the same branch, and a branch always refers to the same series of commits.  In 1.4, branches are modeled similar to Git's tags; they can be created, deleted, and renamed indepedently of commits.

- Simplified commit states.  In prior versions, commits can be in many states including `started`, `finished`, `cancelled`, and `archived`.  In particular, `cancelled` and `archived` have confusing semantics that routinely trip up users.  In 1.4, `cancelled` and `archived` have been removed.

- Flexible pipeline updates.  In prior versions, pipeline updates are all-or-nothing.  That is, an updated pipeline either processes all commits from scratch, or it processes only new commits.  In 1.4, it's possible to have the updated pipeline start processing from any given commit.

- Reduced cluster resource consumption.  In prior versions, each Pachyderm job spawns up a Kubernetes job which in turn spawns up N pods, where N is the user-specified parallelism.  In 1.4, all jobs from a pipeline share N pods.  As a result, a cluster running 1.4 will likely spawn up way fewer pods and use fewer resources in total.

- Simplified deployment dependencies.  In prior versions, Pachyderm depends on RethinkDB and etcd to function.  In 1.4, Pachyderm no longer depends on RethinkDB.

- Dynamic volume provisioning.  GCE and AWS users (Azure support is coming soon) no longer have to manually provision persistent volumes for deploying Pachyderm.  `pachctl deploy` is now able to dynamically provision persistent volumes.  See the [deployment doc](http://pachyderm.readthedocs.io/en/stable/deployment/deploy_intro.html) for details.

Removed features:

A handful of APIs have been removed because they no longer make sense in 1.4.  They include:

- ForkCommit (no longer necessary given the new branch APIs)
- ArchiveCommit (the `archived` commit state has been removed)
- ArchiveAll (same as above)
- DeleteCommit (the original implementation of DeleteCommit is very limiting: only open head commits may be removed.  An improved version of DeleteCommit is coming soon)
- SquashCommit (was only necessary due to the way PPS worked in prior versions)
- ReplayCommit (same as above)

## 1.3.0

Features:

- Embedded Applications - Our “service” enhancement allows you to embed applications, like Jupyter, dashboards, etc., within Pachyderm, access versioned data from within the applications, and expose the applications externally.
- Pre-Fetched Input Data - End-to-end performance of typical Pachyderm pipelines will see a many-fold speed up thanks to a prefetch of input data.
- Put Files via Object Store URLs - You can now use “put-file” with s3://, gcs://, and as:// URLS.
- Update your Pipeline code easily - You can now call “create-pipeline” or “update-pipeline” with the “--push-images” flag to re-run your pipeline on the same data with new images.
- Support for all Docker images - It is no longer necessary to include anything Pachyderm specific in your custom Docker images, so use any Docker image you like (with a couple very small caveats discussed below).
- Cloud Deployment with a single command for Amazon / Google / Microsoft / a local cluster - via `pachctl deploy ...` 
- Migration support for all Pachyderm data from version `1.2.2` through latest `1.3.0`
- High Availability upgrade to rethink, which is now deployed as a petset
- Upgraded fault tolerance via a new PPS job subscription model
- Removed redundancy in log messages, making logs substantially smaller 
- Garbage collect completed jobs
- Support for deleting a commit
- Added user metrics (and an opt out mechanism) to anonymously track usage, so we can discover new bottlenecks
- Upgrade to k8s 1.4.6

## 1.2.0

Features:

- PFS has been rewritten to be more reliable and optimizeable
- PFS now has a much simpler name scheme for commits (eg `master/10`)
- PFS now supports merging, there are 2 types of merge. Squash and Replay
- Caching has been added to several of the higher cost parts of PFS
- UpdatePipeline, which allows you to modify an existing pipeline
- Transforms now have an Env section for specifying environment variables
- ArchiveCommit, which allows you to make commits not visible in ListCommit but still present and readable
- ArchiveAll, which archives all data
- PutFile can now take a URL in place of a local file, put multiple files and start/finish its own commits
- Incremental Pipelines now allow more control over what data is shown
- `pachctl deploy` is now the recommended way to deploy a cluster
- `pachctl port-forward` should be a much more reliable way to get your local machine talking to pachd
- `pachctl mount` will recover if it loses and regains contact with pachd
- `pachctl unmount` has been added, it can be used to unmount a single mount or all of them with `-a`
- Benchmarks have been added
- pprof support has been added to pachd
- Parallelization can now be set as a factor of cluster size
- `pachctl put-file` has 2 new flags `-c` and `-i` that make it more usable
- Minikube is now the recommended way to deploy locally

Content:

- Our developer portal is now available at: http://pachyderm.readthedocs.io/en/latest/
- We've added a quick way for people to reach us on Slack at: http://slack.pachyderm.io
- OpenCV example

## 1.1.0

Features:

- Data Provenance, which tracks the flow of data as it's analyzed
- FlushCommit, which tracks commits forward downstream results computed from them
- DeleteAll, which restores the cluster to factory settings
- More featureful data partitioning (map, reduce and global methods)
- Explicit incrementality
- Better support for dynamic membership (nodes leaving and entering the cluster)
- Commit IDs are now present as env vars for jobs
- Deletes and reads now work during job execution
- pachctl inspect-* now returns much more information about the inspected objects
- PipelineInfos now contain a count of job outcomes for the pipeline
- Fixes to pachyderm and bazil.org/fuse to support writing a larger number of files
- Jobs now report their end times as well as their start times
- Jobs have a pulling state for when the container is being pulled
- Put-file now accepts a -f flag for easier puts
- Cluster restarts now work, even if kubernetes is restarted as well
- Support for json and binary delimiters in data chunking
- Manifests now reference specific pachyderm container version making deployment more bulletproof
- Readiness checks for pachd which makes deployment more bulletproof
- Kubernetes jobs are now created in the same namespace pachd is deployed in
- Support for pipeline DAGs that aren't transitive reductions.
- Appending to files now works in jobs, from shell scripts you can do `>>`
- Network traffic is reduced with object stores by taking advantage of content addressability
- Transforms now have a `Debug` field which turns on debug logging for the job
- Pachctl can now be installed via Homebrew on macOS or apt on Ubuntu
- ListJob now orders jobs by creation time
- Openshift Origin is now supported as a deployment platform

Content:

- Webscraper example
- Neural net example with Tensor Flow
- Wordcount example

Bug fixes:

- False positive on running pipelines
- Makefile bulletproofing to make sure things are installed when they're needed
- Races within the FUSE driver
- In 1.0 it was possible to get duplicate job ids which, that should be fixed now
- Pipelines could get stuck in the pulling state after being recreated several times
- Map jobs no longer return when sharded unless the files are actually empty
- The fuse driver could encounter a bounds error during execution, no longer
- Pipelines no longer get stuck in restarting state when the cluster is restarted
- Failed jobs were being marked failed too early resulting in a race condition
- Jobs could get stuck in running when they had failed
- Pachd could panic due to membership changes
- Starting a commit with a nonexistant parent now errors instead of silently failing
- Previously pachd nodes would crash when deleting a watched repo
- Jobs now get recreated if you delete and recreate a pipeline
- Getting files from non existant commits gives a nicer error message
- RunPipeline would fail to create a new job if the pipeline had already run
- FUSE no longer chokes if a commit is closed after the mount happened
- GCE/AWS backends have been made a lot more reliable

Tests:

From 1.0.0 to 1.1.0 we've gone from 70 tests to 120, a 71% increase.

## 1.0.0 (5/4/2016)

1.0.0 is the first generally available release of Pachyderm.
It's a complete rewrite of the 0.* series of releases, sharing no code with them.
The following major architectural changes have happened since 0.*:

- All network communication and serialization is done using protocol buffers and GRPC.
- BTRFS has been removed, instead build on object storage, s3 and GCS are currently supported.
- Everything in Pachyderm is now scheduled on Kubernetes, this includes Pachyderm services and user jobs.
- We now have several access methods, you can use `pachctl` from the command line, our go client within your own code and the FUSE filesystem layer
