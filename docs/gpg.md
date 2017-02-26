
# GPG

When the service is running it will use it's own key chain location by
internally setting the `GNUPGHOME` environment variable. This will be set
using the config value `duply.config_home`/gpg. It can be confusing
when messing around with keys manually on the host.

Also note that using 8, 16 or 32 characters in key IDs works perfectly fine
in `config.yaml`.

**Use gpg 2.x when possible. 1.x keys are not always compatible.**

## The purpose of the gpg keys

- Each host should have a unique public/private key
- Supply the same public master key to every single host
- Each host will trust the master key meaning you will be able to decrypt
  all backups using master private key you have stored somewhere safe
- If a host is compromised, the intruder will only be able to access backup for the host

## Generating Master Key

Choose RSA 2048. Make sure you give the key a proper user ID so you can locate
the key in your key chain in the future with `gpg --list-keys`.
```
$> gpg --gen-key
...
pub   2048R/6ACEC832 2017-02-19
      Key fingerprint = 05E4 2ABF 5BBD 5D2A 3318  545E FCF8 16B2 6ACE C832
uid                  Your Name (Duply Master Key) <your@email>
sub   2048R/B32D0AE0 2017-02-19
```

We now have master key pair with ID 6ACEC832

### Export
```
$> gpg --armor --export-secret-key -a 6ACEC832 > master_private.key
$> gpg --armor --export -a 6ACEC832 > master_public.key
```

Make sure you store the mater key somewhere safe as the master key can be
used to decrypt all backups (and of course remember the password)

## Generating Public Key

Same procedure as above. Just make sure you are clear about the identity
so you can easily locate keys in the future with `gpg --list-keys`.
