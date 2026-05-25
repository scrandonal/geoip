# GeoIP 增强版：自由定制多种格式 GeoIP 文件

<div align="center">

<a href="https://trendshift.io/repositories/5833" target="_blank"><img src="https://trendshift.io/api/badge/repositories/5833" alt="Loyalsoldier%2Fgeoip | Trendshift" style="width: 250px; height: 55px;" width="250" height="55"/></a>

</div>

<div align="center">
<a href="https://deepwiki.com/Loyalsoldier/geoip" target="_blank"><img src="https://deepwiki.com/badge.svg" alt="Ask DeepWiki badge"></a> <a href="https://www.jsdelivr.com/package/gh/Loyalsoldier/geoip" target="_blank"><img src="https://data.jsdelivr.com/v1/package/gh/Loyalsoldier/geoip/badge?style=rounded" alt="jsdelivr stats badge"></a>

<a href="https://shields.io" target="_blank"><img src="https://img.shields.io/github/downloads/Loyalsoldier/geoip/total?logo=github" alt="GitHub Downloads badge (all assets, all releases)"></a> <a href="https://shields.io" target="_blank"><img src="https://img.shields.io/github/downloads/Loyalsoldier/geoip/latest/total?logo=github" alt="GitHub Downloads badge (all assets, latest release)"></a>
</div>

<p align="center">
  <img src="./assets/hero.png" alt="GeoIP project hero image">
</p>

> **Personal fork note:** This is my personal fork of [Loyalsoldier/geoip](https://github.com/Loyalsoldier/geoip) for learning and local customization. For the official project, please visit the upstream repository.

本项目每周四自动生成多种格式 GeoIP 文件，同时提供命令行界面（CLI）工具供用户自行定制 GeoIP 文件，包括但不限于 V2Ray `dat` 格式文件 `geoip.dat`、MaxMind `mmdb` 格式文件 `Country.mmdb`、sing-box `SRS` 格式文件、mihomo `MRS` 格式文件、Clash ruleset 和 Surge ruleset。

This project releases various formats of GeoIP files automatically every Thursday, and provides a command line interface(CLI) tool for users to customize their own GeoIP files, including but not limited to V2Ray `dat` format file `geoip.dat`, MaxMind `mmdb` format file `Country.mmdb`, sing-box `SRS` format files, mihomo `MRS` format files, Clash ruleset files and Surge ruleset files.

## 与 MaxMind 官方 GeoIP 数据的区别

本项目默认使用 [MaxMind GeoLite2 Country CSV 数据](https://github.com/Loyalsoldier/geoip/blob/release/GeoLite2-Country-CSV.zip)生成各个国家和地区的 GeoIP 文件。所有可供使用的国家和地区 geoip 类别（如 `geoip:cn`，两位英文字母表示国家和地区），请查看：[https://www.iban.com/country-codes](https://www.iban.com/country-codes)。

另外，本项目对 MaxMind 官方 GeoIP 数据做了修改和新增：

- 中国大陆 IPv4 地址数据使用 [@gaoyifan/china-operator-ip](https://github.com/gaoyifan/china-operator-ip/blob/ip-lists/china.txt)
- 中国大陆 IPv6 地址数据使用 [@gaoyifan/china-operator-ip](https://github.com/gaoyifan/china-operator-ip/blob/ip-lists/china6.txt)
- 新增类别（方便有特殊需求的用户使用）：
  - `geoip:cloudflare`（`GEOIP,CLOUDFLARE`）
  - `geoip:cloudfront`（`GEOIP,CLOUDFRONT`）
  - `geoip:facebook`（`GEOIP,FACEBOOK`）
  - `geoip:fastly`（`GEOIP,FASTLY`）
  - `geoip:google`（`GEOIP,GOOGLE`）
  - `geoip:netflix`（`GEOIP,NETFLIX`）
  - `geoip:telegram`（`GEOIP,TELEGRAM`）
  - `geoip:twitter`（`GEOIP,TWITTER`）
  - `geoip:tor`（`GEOIP,TOR`）

## 下载地址与使用方法

本项目发布的所有 GeoIP 文件，请查看 [release 分支](https://github.com/Loyalsoldier/geoip/tree/release)。以下是部分格式 GeoIP 文件的下载地址：

> 如果无法访问域名 `raw.githubusercontent.com`，可以使用第二个地址 `cdn.jsdelivr.net`。
> 如果无法访问域名 `cdn.jsdelivr.net`，可以将其替换为 `fastly.jsdelivr.net`。
>
> *.sha256sum 为校验文件。

### V2Ray dat 格式文件

> 适用于 [V