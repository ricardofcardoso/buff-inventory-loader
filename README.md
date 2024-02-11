# buff-inventory-loader
App to fetch all inventory items from BUFF163 and save them into a MongoDB collection

## Setup
Build the image:

```bash
docker build . -t buff-inventory-loader:v1.0
```

Run it with your environment variables:

```bash
docker run -d --net=host \
-e SESSION_COOKIE=<YOUR_BUFF163_SESSION_COOKIE> \
-e MONGODB_URI=<YOUR_MONGODB_URI> \
--name buff-inventory-loader-container buff-inventory-loader:v1.0
```