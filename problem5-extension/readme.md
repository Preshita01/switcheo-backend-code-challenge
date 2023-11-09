# crude
**crude** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Disclaimer
This solution is **not** based of the scaffold code provided as there were unsolvable compatibility errors when attemping to start a chain.

## Overview
This solution is a possible extension to the current solution shown in *problem5*, where a resource type is first created so that more customised CRUD and query operations can be created. However, it is a preliminary attempt as of now as more time is required to properly implement this solution.

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### CLI Commands

The Crude Module provides the following CLI commands:

- `create-resource`: Create a new resource with the following fields: name (string), category (string), colour (string), size(string) and quantity (uint)
- `update-resouce`: Updates the details of an existing resource based on the newly given details.
- `delete-resource`: Delete a resource using it's ID.
- `list-resource`: List all the available resources.
- `list-resource-colour`: List all the available resources of a specific colour.
- `list-resource-size`: List all the available resources of a specific size.
- `list-resource-quantity`: List all the available resources of a specific quantity.
- `show-resource`: Show the resource with the given ID.
