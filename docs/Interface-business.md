## RESTful Business Module API document



## 1 Get Businesses Interface 

### 1.1 Interface Description    

- Get business list
- Only valid business would be returned

### 1.2 Address  

`{apiAddress}/api/business/get-businesses`

### 1.3 Request Type  

**POST**  

### 1.4 Request Parameters  

#### 1.4.1 Header Parameters  

| Key       | Must | Type/Value      | Note         |
| ------------ | ---- | ---------------- | ------------ |
| Content-Type | Yes   | application/json | Request parameter type |

#### 1.4.2 Body Parameters  

| Key    | Must | Type   | Limit        | Note     |
| --------- | ---- | ------ | --------------- | -------- |
| location_lat   | No | float | 4 digit fraction |  lattitude of the guest's location  |
| location_lng   | No | float | 4 digit fraction |  longitude of the guest's location  |
| sortBy   | Yes | string | 1 < length < 50 | criteria Body parameters should be sorted by |
| startNum | Yes | int | 1 < value | start index of the business in the sorted list |
| quantity   | Yes | int | startNum < value | end index of the business in the sorted list |


**Special Note**:
1. address, zipcode and location info(includes location_lat and location_lng) can't be all empty at the same time
2. check of emptyness should all be done in frontend logic
3. use could search by address or zipcode, but all these has to go through google api and being converted into latitude and longitude pairs
4. we use bid to mark a business. As a result, there should be a table in db to support the mappping from uid to bid.
​    

### 1.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "Get business list successful",  // 提示信息
    "data": [
        {
            "bid": "2312f12dab003e0e",
            "business_name": "foodiePath",
            "business_address": "635 Lexington Ave",
            "business_phoneNum": "4699559587",
            "business_location": {
                "lat": 37.7749,
                "lng": -122.4194 //San Francisco, CA
            },
            "business_availible_time": [{
                "WeekDays": ["Mon", "Tue", "Wen"],
                "Time": "9:00-20:00"
            }],
            "is_available": 1,
            "reason": ""
        },
        {
            "bid": "25411b45452abc76f",
            "business_name": "bestieShop",
            "business_address": "635 3rd Ave",
            "business_phoneNum": "4699165558",
            "business_location": {
                "lat": 37.7749,
                "lng": -122.4194 //San Francisco, CA
            },
            "business_availible_time": [{
                "WeekDays": ["Mon", "Wen"],
                "Time": "9:00-17:00"
            }],
            "is_available": 0,
            "reason": "host is ill today"
        }
    ] 
}
```

### 1.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  




## 2 Get Business Food Interface 

### 2.1 Interface Description    

- Get the food provide by a business

### 2.2 Address  

`{apiAddress}/api/business/get-businesses-food`

### 2.3 Request Type  

**POST**  

### 2.4 Request Parameters  

#### 2.4.1 Header Parameters  

| Key       | Must | Type/Value      | Note         |
| ------------ | ---- | ---------------- | ------------ |
| Content-Type | Yes   | application/json | Request parameter type |

#### 2.4.2 Body Parameters  

| Key    | Must | Type   | Limit        | Note     |
| --------- | ---- | ------ | --------------- | -------- |
| bid   | Yes | string | 1 < length < 20 | business id to locate a business |
| startNum | Yes | int | 1 < value | start index of the food in the sorted list |
| quantity   | Yes | int | startNum < value | end index of the food in the sorted list |


**Special Note**:
1. We could add uid as an input here, once we could do personalized recommendation
2. We cache all the in-day food in Cassandra
3. 

### 2.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "Get business list successful",  // 提示信息
    "data": [
        {
            "bid": "2312f12dab003e0e",
            "food_name": "foodiePathPie",
            "food_ingradients": ["apple", "power","eggs"],
            "food_notes": "Clients may get allergic against blablabla",
            "food_order_cut_time": "20230501T23:59:59",
            "is_permanent": 1,
            "pic": "<picurl>"
        },
        {
            "bid": "25411b45452abc76f",
            "food_name": "bestieShopBowl",
            "food_ingradients": ["white rice", "cucumber","mushroom"],
            "food_notes": "Much carbon, may cause food coma",
            "food_order_cut_time": "202300501T21:59:59",
            "is_permanent": 0,
            "pic": "<picurl>"
        }
    ] 
}
```

### 2.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  
