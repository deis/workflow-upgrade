
|![](https://upload.wikimedia.org/wikipedia/commons/thumb/1/17/Warning.svg/156px-Warning.svg.png) | Deis Workflow is no longer maintained.<br />Please [read the announcement](https://deis.com/blog/2017/deis-workflow-final-release/) for more detail. |
|---:|---|
| 09/07/2017 | Deis Workflow [v2.18][] final release before entering maintenance mode |
| 03/01/2018 | End of Workflow maintenance: critical patches no longer merged |
| | [Hephy](https://github.com/teamhephy/workflow) is a fork of Workflow that is actively developed and accepts code contributions. |

# Workflow Upgrade
[![Docker Repository on Quay](https://quay.io/repository/deisci/workflow-upgrade/status "Docker Repository on Quay")](https://quay.io/repository/deisci/workflow-upgrade)

Deis (pronounced DAY-iss) Workflow is an open source Platform as a Service (PaaS) that adds a developer-friendly layer to any [Kubernetes](http://kubernetes.io) cluster, making it easy to deploy and manage applications on your own servers.

For more information about the Deis Workflow, please visit the main project page at https://github.com/deis/workflow.

We welcome your input! If you have feedback, please [submit an issue][issues]. If you'd like to participate in development, please read the "Development" section below and [submit a pull request][prs].

# About
The Workflow Upgrade service checks for the mismatch between the version of the daemonset pods and the [daemonset][daemons] and deletes the pods such that new pods come up with the latest changes. This is necessary until kubernetes has implemented the update strategy which is tracked through https://github.com/kubernetes/kubernetes/issues/22543.

Workflow Upgrade will be run as kubernetes job after a successful upgrade of the workflow to a new release version.

[issues]: https://github.com/deis/workflow/issues
[prs]: https://github.com/deis/workflow/pulls
[daemons]: http://kubernetes.io/docs/admin/daemons/
[v2.18]: https://github.com/deis/workflow/releases/tag/v2.18.0
