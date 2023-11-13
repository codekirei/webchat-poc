WIP.

Planning to try out a few new-to-me libraries to see how they feel in a
real-ish project.

Will document as I go.

## Development

Start the backend server with `go run . serve`.

## CLI Commands

CLI commands are managed with [Cobra](https://cobra.dev/). Useful docs:

- https://cobra.dev/
- https://github.com/spf13/cobra-cli/tree/main#readme

## Opportunities to Productionize

This is a proof-of-concept, so there are many features that would be present in
a production application have intentionally been omitted. Many such features
are more ops-focused and not in the application layer.

To take this project the extra mile and make it more production-ready, here are
some things to add:

- proper migration handling for database changes
- proper ci/cd
- swagger api documentation
- data backups
- compliance with data retention regulation (i.e. support user data deletion
  and export for [GDPR](https://gdpr.eu/))
- a more performant DB for the data model (specifically chat history --
  [Cassandra](https://cassandra.apache.org/_/index.html) is a good option)
- structured logging with external aggregation
  (i.e.[zerolog](https://github.com/rs/zerolog) and
  [Loki](https://grafana.com/docs/loki/latest/) or
  [openobserve](https://github.com/openobserve/openobserve) or
  [highlight](https://www.highlight.io/), with [vector](https://vector.dev/) or
  [fluentbit](https://fluentbit.io/) if data transformation is needed)
- metrics, monitoring, and alerts (i.e. [Grafana](https://grafana.com/) and
  [Prometheus](https://prometheus.io/))
- error tracking via an external service (i.e. [Sentry](https://sentry.io/) or
  [GlitchTip](https://glitchtip.com/))
