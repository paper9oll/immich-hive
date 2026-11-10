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

Create the config directory and download the example configuration:

```sh
mkdir -p config
curl -fsSL https://raw.githubusercontent.com/paper9oll/immich-hive/main/config.yml.example -o config/config.yml
```

### 2. Configure `config.yml`

#### Immich API Key

1. Go to your Immich server (for example, `http://immich-server:2283`).
2. Go to **Account Settings** > **API Keys**.
3. Click **New API Key**.
4. Check the following permissions:

   <details>
   <summary>List of required permissions</summary>

   **Asset**
   - `asset.read`
   - `asset.view`
   - `asset.download`

   **Album**
   - `album.read`
   - `album.statistics`

   **Face**
   - `face.read`

   **Person**
   - `person.read`
   - `person.statistics`

   **Server**
   - `server.about`

   **Tag**
   - `tag.read`
   </details>

5. Click **Create**, then **Copy to Clipboard**.
6. Paste the API Key into `immich_api_key`.

#### Weather

1. Log in to the [OpenWeatherMap API keys page](https://home.openweathermap.org/api_keys).
2. Create a new API key or copy the **Default** API key.
3. Paste the API key into `weather.locations.api`.
4. Enter the location name in `weather.locations.name`.
5. Enter the location's latitude in `weather.locations.lat`.
6. Enter the location's longitude in `weather.locations.lon`.
7. Optionally, set `weather.locations.unit` to `standard` or `imperial`.
8. Optionally, set `weather.locations.lang` to your preferred language. See OpenWeatherMap's [supported languages](https://openweathermap.org/api/current?collection=current_forecast#multi) for details.

<details>
<summary>Example</summary>

```yml
weather:
  locations:
    - name: 'London'
      lat: '51.5080'
      lon: '-0.1281'
      api: 'YOUR_API_KEY'
      unit: metric
      round_temperature: true
      lang: en
```

</details>

#### Other settings

For settings not covered above, refer to the [Immich Kiosk documentation](https://docs.immichkiosk.app/configuration/).

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

OpenWeatherMap is a trademark of OpenWeather Ltd. This project is not affiliated with, endorsed by, or sponsored by OpenWeather Ltd.
