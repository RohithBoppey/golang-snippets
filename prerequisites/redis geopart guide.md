Your key is:

```
orders:locations
```

Below is a **redis-cli cheat sheet**, mapped directly to what your Go code is doing + common debugging queries you’ll actually need.

---

## 🔑 Key basics

### Check if the key exists

```bash
EXISTS orders:locations
```

---

## 📍 How many points (orders) are in DB

### Total number of geo points

```bash
ZCARD orders:locations
```

> This matches: _“how many orders are stored right now?”_

---

## ➕ Add a geo point (same as `GeoAdd`)

```bash
GEOADD orders:locations <longitude> <latitude> <order_id>
```

Example:

```bash
GEOADD orders:locations 77.5946 12.9716 order_123
```

---

## ❌ Delete a single entry (same as `Remove(orderID)`)

```bash
ZREM orders:locations <order_id>
```

Example:

```bash
ZREM orders:locations order_123
```

---

## 🧹 Delete all entries (same as `RemoveAll()`)

### Delete only this key

```bash
DEL orders:locations
```

### (Optional) Delete all keys in Redis (⚠️ dangerous)

```bash
FLUSHALL
```

---

## 📋 List all order IDs (members)

```bash
ZRANGE orders:locations 0 -1
```

> This gives **only IDs**, no lat/lon.

---

## 🌍 Get coordinates of an order

```bash
GEOPOS orders:locations <order_id>
```

Example:

```bash
GEOPOS orders:locations order_123
```

Output:

```
1) 1) "77.5946"
   2) "12.9716"
```

---

## 📏 Distance between two orders

```bash
GEODIST orders:locations <order_id_1> <order_id_2> km
```

Example:

```bash
GEODIST orders:locations order_123 order_456 km
```

---

## 🔎 Nearby search (matches `GetNearbyToLocation`)

### Find nearby orders within radius (km)

```bash
GEOSEARCH orders:locations
  FROMLONLAT <lon> <lat>
  BYRADIUS <radius> km
```

Example:

```bash
GEOSEARCH orders:locations \
  FROMLONLAT 77.5946 12.9716 \
  BYRADIUS 5 km
```

---

## 🔎 Nearby search with distance + coordinates

(matches `WithDist: true`, `WithCoord: true`)

```bash
GEOSEARCH orders:locations \
  FROMLONLAT <lon> <lat> \
  BYRADIUS <radius> km \
  WITHDIST \
  WITHCOORD \
  COUNT 10 \
  ASC
```

Example:

```bash
GEOSEARCH orders:locations \
  FROMLONLAT 77.5946 12.9716 \
  BYRADIUS 10 km \
  WITHDIST \
  WITHCOORD \
  COUNT 10 \
  ASC
```

Output structure:

```
1) 1) "order_123"
   2) "1.42"
   3) 1) "77.5946"
      2) "12.9716"
```

---

## 🧠 Debugging / inspection queries

### See raw sorted-set scores (geohashes)

```bash
ZRANGE orders:locations 0 -1 WITHSCORES
```

### Memory used by the key

```bash
MEMORY USAGE orders:locations
```

### Type check (should be zset)

```bash
TYPE orders:locations
```

---

## 🚨 Common mistakes to avoid

- ❌ `GET orders:locations` → ❌ wrong (not a string)
- ❌ `HGETALL orders:locations` → ❌ wrong (not a hash)
- ✅ Always use `Z*` or `GEO*` commands

---

## 🧩 Mental model (important)

What your code _conceptually_ does:

| Go code      | Redis reality            |
| ------------ | ------------------------ |
| `GeoAdd`     | `ZADD` with geohash      |
| `GeoSearch`  | `ZRANGE` + geo filtering |
| `Remove`     | `ZREM`                   |
| Count orders | `ZCARD`                  |

---

If you want, next I can:

- give you **prod-safe debugging commands**
- help you add **TTL / expiry per order**
- show **how to store order metadata alongside GEO**

Just tell me 👍
