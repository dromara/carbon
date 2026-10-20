# Carbon docs

Documentation for Carbon, built with VitePress.

## Development

Install the dependencies and start the VitePress development server:

```bash
yarn install
yarn docs:dev
```

## Build

Build the documentation site locally with:

```bash
yarn docs:build
```

The generated files are written to the `dist` directory.

## Deployment

The documentation is deployed through GitHub Pages Actions. The workflow in
`.github/workflows/deploy-docs.yml` runs when changes are pushed to
the `docs` branch or when it is triggered manually. It builds the VitePress
site and publishes the generated `dist` directory to GitHub Pages.

Website: https://dromara.github.io/carbon
