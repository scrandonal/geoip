# 配置说明

## 总览

本项目使用的配置文件格式为 `json`，其中包含 `input` 和 `output` 两个数组，每个数组包含一个或多个输入或输出格式的具体配置。

```json
{
  "input":  [],
  "output": []
}
```

## 支持的输入或输出格式

支持的 `input` 输入格式：

- **clashRuleSet**：ipcidr 类型的 Clash RuleSet
- **clashRuleSetClassical**：classical 类型的 Clash RuleSet
- **cutter**：用于裁剪前置步骤中的数据
- **dbipCountryMMDB**：DB-IP country mmdb 数据格式（`dbip-country-lite.mmdb`）
- **ipinfoCountryMMDB**：IPInfo country mmdb 数据格式（`country.mmdb`）
- **json**：JSON 数据格式
- **maxmindGeoLite2ASNCSV**：MaxMind GeoLite2 ASN CSV 数据格式（`GeoLite2-ASN-CSV.zip`）
- **maxmindGeoLite2CountryCSV**：MaxMind GeoLite2 country CSV 数据格式（`GeoLite2-Country-CSV.zip`）
- **maxmindMMDB**：MaxMind GeoLite2 country mmdb 数据格式（`GeoLite2-Country.mmdb`）
- **mihomoMRS**：mihomo MRS 数据格式（`geoip-cn.mrs`）
- **private**：局域网和私有网络 CIDR（例如：`192.168.0.0/16` 和 `127.0.0.0/8`）
- **singboxSRS**：sing-box SRS 数据格式（`geoip-cn.srs`）
- **stdin**：从 standard input 获取纯文本 IP 和 CIDR（例如：`1.1.1.1` 或 `1.0.0.0/24`）
- **surgeRuleSet**：Surge RuleSet
- **text**：纯文本 IP 和 CIDR（例如：`1.1.1.1` 或 `1.0.0.0/24`）
- **v2rayGeoIPDat**：V2Ray GeoIP dat 数据格式（`geoip.dat`）

支持的 `output` 输出格式：

- **clashRuleSet**：ipcidr 类型的 Clash RuleSet
- **clashRuleSetClassical**：classical 类型的 Clash RuleSet
- **dbipCountryMMDB**：DB-IP country mmdb 数据格式（`dbip-country-lite.mmdb`）
- **ipinfoCountryMMDB**：IPInfo country mmdb 数据格式（`country.mmdb`）
- **lookup**：从指定的列表中查找指定的 IP 或 CIDR
- **maxmindMMDB**：MaxMind GeoLite2 country mmdb 数据格式（`GeoLite2-Country.mmdb`）
- **mihomoMRS**：mihomo MRS 数据格式（`geoip-cn.mrs`）
- **singboxSRS**：sing-box SRS 数据格式（`geoip-cn.srs`）
- **stdout**：将纯文本 CIDR 输出到 standard output（例如：`1.0.0.0/24`）
- **surgeRuleSet**：Surge RuleSet
- **text**：纯文本 CIDR（例如：`1.0.0.0/24`）
- **v2rayGeoIPDat**：V2Ray GeoIP dat 数据格式（`geoip.dat`）

## `input` 输入格式配置项

### **clashRuleSet**

- **type**：（必须）输入格式的名称
- **action**：（必须）操作类型，值为 `add`（添加 IP 地址）或 `remove`（移除 IP 地址）
- **args**：（必须）
  - **name**：类别名称。（不能与 `inputDir` 同时使用；需要与 `uri` 同时使用）
  - **uri**：Clash `ipcidr` 类型的 ruleset 文件路径，可为本地文件路径或远程 `http`、`https` 文件 URL。（不能与 `inputDir` 同时使用；需要与 `name` 同时使用）
  - **inputDir**：需要遍历的输入目录（不遍历子目录）。（遍历的文件名作为类别名称；不能与 `name` 和 `uri` 同时使用）
  - **wantedList**：（可选，数组）指定需要的类别/文件。（与 `inputDir` 同时使用）
  - **onlyIPType**：（可选）只处理的 IP 地址类型，值为 `ipv4` 或 `ipv6`。

> **个人备注**：日常使用时建议优先用 `onlyIPType: "ipv4"` 以减少规则集体积，IPv6 按需单独生成。

```jsonc
{
  "type": "clashRuleSet",
  "action": "add",     // 添加 IP 地址
  "args": {
    "name": "cn",
    "uri": "./cn.yaml" // 读取本地文件 cn.yaml 的 IPv4 和 IPv6 地址，并添加到 cn 类别中
  }
}
```

```jsonc
{
  "type": "clashRuleSet",
  "action": "add",                    // 添加 IP 地址
  "args": {
    "inputDir": "./clash/yaml",       // 遍历 ./clash/yaml 目录内的所有文件（不遍历子目录）
    "wantedList": ["cn", "us", "jp"], // 只需要 ./clash/yaml 目录内文件名去除扩展名后，名为 cn、us、jp 的文件
    "onlyIPType": "ipv6"              // 只添加 IPv6 地址
  }
}
```

```jsonc
{
  "type": "clashRuleSet",
  "action": "remove",                     // 移除 IP 地址
  "args": {
    "name": "cn",
    "uri": "https://example.com/cn.yaml", // 读取网络文件内容
    // 注意：从远程 URI 拉取时建议在 CI 中设置超时，避免因网络问题导致构建卡死
    "onlyIPType": "ipv4"                  // 只移除 IPv4 地址
  }
}
```
