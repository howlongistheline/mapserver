# mapserver
problem: cannot easily render a map with pin on coordinates


### API
This is now running at https://howlongistheline.org/maps

Simply request lattitude,longitude,API-key and it will return a map tile with a pin at the coordinates provided.

If you are going to be using this for something other than howlongistheline.org and expect a non-trivial number of requests you should request an API key. Otherwise use the public key: `K3340`

#### Example usage:
`curl https://howlongistheline.org/maps/52.512411,13.381612,K3340 -o test.png`

or simply open [https://howlongistheline.org/maps/52.512411,13.381612,K3340](https://howlongistheline.org/maps/52.512411,13.381612,K3340) in your browser.
