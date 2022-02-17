# Typescript Monorepo

In this example we are in a typescript monorepo using yarn workspaces.

## Workspaces

In [package.json](./package.json) our workspaces are defined as the following:

```json
{
	// ...
	"private": "true",
	"workspaces": ["packages/*", "apps/*"]
	// ...
}
```

This will allow yarn to handle all our packages in `packages` and `apps` from the root of our project with the `yarn install` command.

For more info about yarn workspaces check the documentation [here](https://classic.yarnpkg.com/lang/en/docs/workspaces/).

## Structure

Our monorepo is a basic example. We just have the [@example/tsconfig](./packages/tsconfig/) package on which all our typescript packages will depend.

It just contains a basic `tsconfig.json` where all our shared typescript options are defined.

## Templays

Our templay in [templays](./templays/) folder will help us setting up new packages in our monorepo with a single command.

This is our [module](./templays/module/) templay in `.templays.yml`:

```yaml
templays:
  module: ./templays/module
```

In our templay we will have a `src` directory with a simple `index.ts` file, a `tsconfig.json` which extends from our [@example/tsconfig](./packages/tsconfig/) package, and a `package.json` file.

The `package.json` file has a templay variable defined as the following:

```json
{
	"name": "@example/{{.module}}"
	// ...
}
```

In this way we will be able to specify our new package's name by passing a templay variable named `module`

So we just defined a templay called `module` defined in our `templays/module` folder.

If we now run the list command

```console
$ yarn templay list
```

we will get the following output:

```
Name       Path
module     ./templays/module
```

This will indicate that the templay is set correctly.

Now we are ready to generate!

## Generate new packages

To generate our new templayed packages we just need to run the following command:

```console
$ yarn templay gen -d destination/path -v module="module name" module
```

In this example we generated `apps/generated1` and `packages/generated2` packages by running:

```console
$ yarn templay gen -d apps/generated1 -v module=generated1 module
```

and

```console
$ yarn templay gen -d packages/generated2 -v module=generated2 module
```

getting the follwing outputs:

```
Templay module successfully generated in apps/generated1
```

and

```
Templay module successfully generated in packages/generated2
```
