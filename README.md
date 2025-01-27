# Dasharr

 Dashboard of your indexers' usage

 <p align="center">
  <a href="https://discord.gg/4vd7qAaFwX">
    <img src="https://img.shields.io/badge/Discord-Chat-5865F2?logo=discord&logoColor=white" alt="Join Our Discord">
  </a>
</p>

 > Note: this project is in very early stages, expect broken stuff and breaking changes !

 ![header](images/0.png)
 ![header](images/1.png)

## Why ?

We often times don't realize the evolution of the stats on torrent indexers, this fixes it. Also :

- see the impact of automated tools
- see the impact of new uploads
- have a clearer view of how much bounty can be spent on requests every n days
- because it's cool

 <details> <summary> <b> Supported indexers </b> </summary>

 * ANT
 * RED
 * GGn
 * OPS
 * BLU
 * Aither
 * ItaTorrents
 * OTW
 * LST
 * seedpool

</details>



 ## Quickstart

 ### Docker

- Copy [docker-compose.yml](./docker-compose.yml), edit the fields you want (most importantly the `API_KEY` env var).
- Run the container `docker-compose up -d`
- Go to the config volume and enable the indexers you want in `config.json`
- Visit the webui and configure the indexers that require configuration
- Restart the container, a first stats collection will be made
- New stats will be collected for the enabled indexers every 6h