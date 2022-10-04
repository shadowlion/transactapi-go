# How to Contribute

It's relatively simple! Take a look at the [documentation](https://api.norcapsecurities.com/admin_v3/documentation), find which endpoint hasn't been covered
yet. Normally, you will want to do the following on github:

1. Open an issue

2. Open a pull request addressing the issue

3. Create a folder in `/src` with the name of the `ENDPOINT`. Create the following file:

   - `/transactapi/ENDPOINT.go` - function file to call the api endpoint

4. Once all the lints/tests pass locally, bump the version number (e.g. 0.0.1 -> 0.0.2).

5. Await the PR to be approved!
