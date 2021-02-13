# kumade

ConoHa API CLI

## config

~/.config/kumade/config.json
```json
{
  "user": "username",
  "password": "password",
  "tenant_id": "tenant_id"
}
```

もしくは

```env
export KUMADE_USERNAME=username
export KUMADE_PASSWORD=password
export KUMADE_TENANT_ID=tenant_id
```

## command

### get token info
```
$ kumade identify token
```

### get image list
```
$ kumade image images
```
