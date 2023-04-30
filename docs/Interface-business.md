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
| location_lat   | Yes | float | 4 digit fraction |  lattitude of the guest's location  |
| location_lng   | Yes | float | 4 digit fraction |  longitude of the guest's location  |
| open | Yes | int | 1 digit | a bit controls whether to only show opening business |
| startNum | Yes | int | 1 < value | start index of the business in the sorted list |
| quantity   | Yes | int | startNum < value | end index of the business in the sorted list |


**Special Note**:
1. We use bid to mark a business. As a result, there should be a table in db to support the mappping from uid to bid.
2. Returned businesses should be sorted ascending with the linear distance with guest.
3. TODO: add sortBy key after we have the business review and user histories.
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


# 2 Get Businesses by Name Interface 

### 2.1 Interface Description    

- Get details of a list of business by name searching
- All related business would be returned no matter it is valid or not

### 2.2 Address  

`{apiAddress}/api/business/get-businesses-by-name`

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
| business-name   | Yes | string | 1 < length < 50  | a piece of business name |
| zipcode   | No | string | 5 digit |  guests' zipcode  |
| city   | Yes | string | 1 < length < 30 |  guests city name  |
| State   | Yes | string | 1 < length < 30 | guests state name |
| open | Yes | int | 1 digit | a bit controls whether to only show opening business |
| startNum | Yes | int | 1 < value | start index of the business in the sorted list |
| quantity   | Yes | int | startNum < value | end index of the business in the sorted list |


**Special Note**:
1. zipcode could used to sort. Once a zipcode given by a user, the businesses returned are sorted by closeness to toward the certer of the zipcode
2. we use bid to mark a business. As a result, there should be a table in db to support the mappping from uid to bid.
3. TODO: add sortBy key after we have the business review and user histories.
​    

### 2.5 Sample Response

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

### 2.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  


## 3 Get Business Food Interface 

### 3.1 Interface Description    

- Get the food provide by a business

### 3.2 Address  

`{apiAddress}/api/business/get-businesses-food`

### 3.3 Request Type  

**POST**  

### 3.4 Request Parameters  

#### 3.4.1 Header Parameters  

| Key       | Must | Type/Value      | Note         |
| ------------ | ---- | ---------------- | ------------ |
| Content-Type | Yes   | application/json | Request parameter type |

#### 3.4.2 Body Parameters  

| Key    | Must | Type   | Limit        | Note     |
| --------- | ---- | ------ | --------------- | -------- |
| bid   | Yes | string | 1 < length < 20 | business id to locate a business |
| startNum | Yes | int | 1 < value | start index of the food in the sorted list |
| quantity   | Yes | int | startNum < value | end index of the food in the sorted list |


**Special Note**:
1. We could add uid as an input here, once we could do personalized recommendation
2. We cache all the in-day food in Cassandra
3. 

### 3.5 Sample Response

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

### 3.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  
