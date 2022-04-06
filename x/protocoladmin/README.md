# Protocoladmin module

This module contain custom authentication functionality
Inorder to start with this module, it needs to have super admin from genesis state.

Geneiss app_state example

```yml
app_state:
    protocoladmin:
      groupList: [
        {
          name: "super.admin",
          owner: "genesis"
        }
      ]
      adminList: [
        {
          group: "super.admin",
          admin: <super.admin.address>
        }
      ]
```

### __command list__

__Query__

```devnet-wasmd query protocoladmin```

```bash
Available Commands:
  list-admin          list all admin
  list-admin-of-group Query list-admin-of-group
  list-group          list all group
  params              shows the parameters of the module
  show-admin          shows a admin
  show-group          shows a group
```

list-admin

```bash
devnet-wasmd query protocoladmin list-admin
```

list-admin-of-group

```bash
devnet-wasmd query protocoladmin list-admin-of-group [group] [flags]
```

list-group

```bash
devnet-wasmd query protocoladmin list-group
```

show-admin

```bash
devnet-wasmd query protocoladmin show-admin [group] [admin]
```

show-group

```bash
devnet-wasmd query protocoladmin show-group [name]
```

__Tx__

```bash
Available Commands:
  add-admin-to-group      Broadcast message add-admin-to-group
  create-group            Create a new group
  delete-group            Delete a group
  remove-admin-from-group Broadcast message remove-admin-from-group
  update-group            Update a group
```

add-admin-to-group

Inorder to add admin to group, the group needs to be existed and msg sender needs to be the group creator or super admin

```bash
devnet-wasmd tx protocoladmin add-admin-to-group [name] [address]
```

create-group

msg sender have to be super admin

```bash
devnet-wasmd tx protocoladmin create-group [name]
```

delete-group

msg sender have to be super admin or group owner

```bash
devnet-wasmd tx protocoladmin delete-group [name]
```

remove-admin-from-group

msg sender have to be super admin or group owner

```bash
devnet-wasmd tx protocoladmin remove-admin-from-group [name] [address]
```

update-group

msg sender have to be super admin or group owner

```bash
devnet-wasmd tx protocoladmin update-group [name]
```
