# Immich Hive

An opinionated version of [Immich Kiosk](https://github.com/damongolding/immich-kiosk) with styling inspired by Google Nest.

## Install

### 1. Download the required files

Create a directory of your choice (for example, `./immich-hive`) to hold `docker-compose.yml`.

```sh
mkdir -p ./immich-hive
cd ./immich-hive
```

Download [docker-compose.yml](https://raw.githubusercontent.com/paper9oll/immich-hive/main/docker-compose.yml) by running:

```sh
curl -fsSL https://raw.githubusercontent.com/paper9oll/immich-hive/main/docker-compose.yml -o docker-compose.yml
```

Alternatively, download the file in your browser and move it to the directory you created.

### 2. Configure `docker-compose.yml` file

Set `KIOSK_IMMICH_API_KEY` to an API key for your Immich server. Refer to the [Immich Kiosk documentation](https://docs.immichkiosk.app/installation/#api-key-permissions) on the required API key permissions.

For the full configuration details, refer to the [Immich Kiosk documentation](https://docs.immichkiosk.app/configuration/).

### 3. Start the container

From the directory you created in step 1, run:

```sh
docker compose up -d
```

## Post-installation

Access the web application at `http://IP-ADDRESS:3000`.

> [!NOTE]
> Replace `IP-ADDRESS` with the IP address of your Docker host.

## License

Immich Hive is licensed under the [GNU Affero General Public License, version 3](LICENSE).

## Trademarks

Immich and its associated marks belong to their respective owner(s). This project is not affiliated with, endorsed by, or sponsored by the Immich project.

Immich Kiosk and its associated marks belong to their respective owner(s). This project is not affiliated with, endorsed by, or sponsored by the Immich Kiosk project.

Google, Google Nest, and Nest are trademarks of Google LLC. This project is not affiliated with, endorsed by, or sponsored by Google LLC.
