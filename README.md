## Insights Client: TNG ##

Client collection is broken into stages:

1. Self update - the client code itself is updated
2. Core update - the core module (usually distributed as an egg hosted on cert-api.access.redhat.com)
3. Collection - the currently downloaded core module is invoked and an archive is created
4. Upload - the archive is submitted to the service for analysis
5. Show - the latest rule matches are retrieved from the service and printed

Each stage is run individually using a subcommand syntax. For example, to run
the self update stage:

```bash
$ insights selfupdate
```

Stages will run and then exit. To run the entire update-collect-upload-show
pipeline, run `insights` without a subcommand.

```bash
$ insights
```

Each subcommand supports a variety of options to adjust behavior. See the
output of `--help` for the subcommand for details.

## Legacy Client Support ##

A wrapper bash script provides an CLI compatibility interface from the old
"insights-client" CLI to the new "insights" CLI.