<img src='doc/pachyderm_factory_gh.png' height='225' title='Pachyderm'> 

[![GitHub release](https://img.shields.io/github/release/pachyderm/pachyderm.svg?style=flat-square)](https://github.com/pachyderm/pachyderm/releases)
[![GitHub license](https://img.shields.io/badge/license-apache-blue.svg)](https://github.com/pachyderm/pachyderm/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/pachyderm/pachyderm?status.svg)](https://godoc.org/github.com/pachyderm/pachyderm/src/client)
[![Go Report Card](https://goreportcard.com/badge/github.com/pachyderm/pachyderm)](https://goreportcard.com/report/github.com/pachyderm/pachyderm)
[![Slack Status](http://slack.pachyderm.io/badge.svg)](http://slack.pachyderm.io)

# Pachyderm: Data Pipelines

Pachyderm is a tool for production data pipelines. If you need to chain
together data scraping, ingestion, cleaning, munging, wrangling, processing,
modelling, and analysis in a sane way, then Pachyderm is for you. If you have an
existing set of scripts which do this in an ad-hoc fashion and you're looking
for a way to "productionize" them, Pachyderm can make this easy for you.

## Features

- Containerized: Pachyderm is built on Docker and Kubernetes. Whatever
  languages or libraries your pipeline needs, they can run on Pachyderm which
  can easily be deployed on any cloud provider or on prem.
- Version Control: Pachyderm version controls your data as it's processed. You
  can always ask the system how data has change, see a diff, and, if something
  doesn't look right, revert.
- Provenance (aka data lineage): Pachyderm tracks where data comes from. Pachyderm keeps track of all the code and  data that created a result.
- Parallelization: Pachyderm can efficiently schedule massively parallel
  workloads.
- Incremental Processing: Pachyderm understands how your data has changed and
  is smart enough to only process the new data.

## Getting Started
[Install Pachyderm locally](http://pachyderm.readthedocs.io/en/latest/getting_started/local_installation.html) or [deploy on AWS/GCE/Azure](http://pachyderm.readthedocs.io/en/latest/deployment/deploy_intro.html) in about 5 minutes. 

You can also refer to our complete [developer docs](http://pachyderm.readthedocs.io/en/latest) to see tutorials, check out example projects, and learn about advanced features of Pachyderm.

If you'd like to see some examples and learn about core use cases for Pachyderm:
- [Examples](http://pachyderm.readthedocs.io/en/latest/examples/readme.html)
- [Use Cases](http://www.pachyderm.io/use_cases.html)
- [Case Studies](http://www.pachyderm.io/usecases/generalfusion.html): Learn how [General Fusion](http://www.generalfusion.com/) uses Pachyderm to power commercial fusion research.

## Documentation

[Official Documentation](http://pachyderm.readthedocs.io/en/latest/)

## Community
Keep up to date and get Pachyderm support via:
- [Twitter](http://twitter.com/pachydermio)
- [![Slack Status](http://slack.pachyderm.io/badge.svg)](http://slack.pachyderm.io) Join our community [Slack Channel](http://slack.pachyderm.io) to get help from the Pachyderm team and other users.

### Contributing

To get started, sign the [Contributor License Agreement](https://pachyderm.wufoo.com/forms/pachyderm-contributor-license-agreement).

You should also check out our [contributing guide](./contributing).

Send us PRs, we would love to see what you do! You can also check our GH issues for things labeled "noob-friendly" as a good place to start. We're sometimes bad about keeping that label up-to-date, so if you don't see any, just let us know. 

### Join Us

WE'RE HIRING! Love Docker, Go and distributed systems? Learn more about [our team](http://www.pachyderm.io/jobs.html) and email us at jobs@pachyderm.io.

## Usage Metrics

Pachyderm automatically reports anonymized usage metrics. These metrics help us
understand how people are using Pachyderm and make it better.  They can be
disabled by setting the env variable `METRICS` to `false` in the pachd
container.
