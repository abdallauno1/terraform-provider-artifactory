---
subcategory: "Security"
---
# Artifactory keypair Resource

RSA key pairs are used to sign and verify the Alpine Linux index files in JFrog Artifactory, while GPG key pairs are 
used to sign and validate packages integrity in JFrog Distribution. The JFrog Platform enables you to manage multiple 
RSA and GPG signing keys through the Keys Management UI and REST API. The JFrog Platform supports managing multiple 
pairs of GPG signing keys to sign packages for authentication of several package types such as Debian, Opkg, and RPM 
through the Keys Management UI and REST API.
Passphrases are not currently supported, though they exist in the API.


## Example Usage

```hcl
terraform {
  required_providers {
    artifactory = {
      source    = "registry.terraform.io/jfrog/artifactory"
      version   = "2.6.14"
    }
  }
}
resource "artifactory_keypair" "some-keypair6543461672124900137" {
  pair_name   = "some-keypair6543461672124900137"
  pair_type   = "RSA"
  alias       = "foo-alias6543461672124900137"
  private_key = file("samples/rsa.priv")
  public_key  = file("samples/rsa.pub")
  
  lifecycle {
    ignore_changes = [
      private_key,
      passphrase,
    ]
  }
}
```

## Argument Reference

The following arguments are supported:

* `pair_name` - (Required) A unique identifier for the Key Pair record.
* `pair_type` - (Required) Key Pair type. Supported types - GPG and RSA.
* `alias` - (Required) Will be used as a filename when retrieving the public key via REST API.
* `private_key` - (Required, Sensitive)  - Private key. PEM format will be validated.
* `passphrase` - (Optional) Passphrase will be used to decrypt the private key. Validated server side.
* `public_key` - (Required) Public key. PEM format will be validated.
* `unavailable` - (Computed) Unknown usage. Returned in the json payload and cannot be set.

Artifactory REST API call Get Key Pair doesn't return keys `private_key` and `passphrase`, but consumes these keys in the POST call.
The meta-argument `lifecycle` used here to make Provider ignore the changes for these two keys in the Terraform state.

## Import

Keypair can be imported using their name, e.g.

```
$ terraform import artifactory_keypair.my-keypair my-keypair
```