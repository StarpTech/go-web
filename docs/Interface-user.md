# RESTful User Module API document  

## 1 User Registration Interface 

### 1.1 Interface Description    

- User info registration  
- User sign up through phone# and email
- User sign up through facebood or google account(TODO)
- phone# and email binding with unique account 

### 1.2 Address  

`{apiAddress}/api/user/signup`  

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
| account   | Yes   | string | 1 < length < 50 | Account |
| passcode  | Yes   | string | 1 < length < 50 | Passord     |
| email  | Yes   | string | 1 < length < 50 | email binded with the account  |
| phoneNum  | Yes   | string | 1 < length < 20 | user phone number     |
| checkCode | Yes   | string | length = 6      | Check Code   |
| waiverFlag | Yes   | int | {1, 0}      | Has user signed wiver   |

**Special Note**: passcode encription with SHA256  (* to be decided)

**Other Interfaces needs to be Involved**:  

| Interface Name   | Interface Address                               | Purpose           |
| ---------- | -------------------------------------- | ------------------ |
| Get Checkcode | `{apiAddress}/api/common/getCheckCode` | getNecessary Checkcode |
| Get Waiver | `{apiAddress}/api/common/getWaiver` | get the Waiver needs to signed by user |

​    

### 1.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "Registration Successful",  // 提示信息
    "data": null  // 返回内容
}
```

### 1.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  

  

      

## 2 User login Interface 

### 2.1 Interface Description    

- User login through account, email or phone# and password 
- User login through facebood or google number (TODO)

### 2.2 Address  

`{apiAddress}/api/user/login`  

### 2.3 Request Type  

**POST**  

### 2.4 Request Parameters  

#### 2.4.1 Header Parameters  

| Key       | Must | Type/Value      | Note         |
| ------------ | ---- | ---------------- | ------------ |
| Content-Type | Yes   | application/json | Request parameter type |

#### 2.4.2 Body Parameters  

| Key    | Mush | Type   | Limit        | Note     |
| --------- | ---- | ------ | --------------- | -------- |
| account   | No   | string | 1 < length < 50 | Account |
| email   | No   | string | 1 < length < 50 | Useer email |
| phoneNum   | No   | string | 1 < length < 50 | User phoneNum |
| passcode  | Yes   | string | 1 < length < 50 | Passord     |
| checkCode | Yes   | string | length = 6      | Check Code   |

**Special Note**: 
1. passcode encription with SHA256  (* to be decided)
2. account, email and phoneNum cannot be None at the same time 
3. Verification of the triple None logic should be put at frontend

**Other Interfaces needs to be Involved**:  

| Interface Name   | Interface Address                               | Purpose           |
| ---------- | -------------------------------------- | ------------------ |
| Get Checkcode | `{apiAddress}/api/common/getCheckCode` | getNecessary Checkcode |

​    

### 2.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "Login Successful",  // 提示信息
    "data": {
        "user_id": "32d23ab3b420ef"
    }  // 返回内容
}
```

​    

### 2.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  




## 3 Set User Question Interface 

### 3.1 Interface Description    

- Set up user password questions and answers  
- Can be used either for add or update
- Set up during registration
- Can be modified after registration 

### 3.2 Address  

`{apiAddress}/api/user/set-question`  

### 3.3 Request Type  

**POST**  

### 3.4 Request Parameters  

#### 3.4.1 Header Parameters  

| Key       | Must | Type/Value      | Note         |
| ------------ | ---- | ---------------- | ------------ |
| Content-Type | Yes   | application/json | Request parameter type |

#### 3.4.2 Body Parameters  

| Key    | Mush | Type   | Limit        | Note     |
| --------- | ---- | ------ | --------------- | -------- |
| question   | Yes   | string | 1 < length < 50 | Passord Question |
| answer  | Yes   | string | 1 < length < 50 | Answer for the Question |
| userId  | Yes   | string | 1 < length < 50 | UserId |


### 3.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "Question set up successfully",  // 提示信息
    "data": null  // 返回内容
}
```

​    

### 3.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  



## 4 Get User Question Interface 

### 4.1 Interface Description    

- Get user password questions and answers  

### 4.2 Address  

`{apiAddress}/api/user/get-question`  

### 4.3 Request Type  

**POST**  

### 4.4 Request Parameters  

#### 4.4.1 Header Parameters  

| Key       | Must | Type/Value      | Note         |
| ------------ | ---- | ---------------- | ------------ |
| Content-Type | Yes   | application/json | Request parameter type |

#### 4.4.2 Body Parameters  

| Key    | Mush | Type   | Limit        | Note     |
| --------- | ---- | ------ | --------------- | -------- |
| userId   | Yes   | string | 1 < length < 50 | Account |
| answer  | Yes   | string | 1 < length < 50 | Passord     |


### 4.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "User Questions are in the data field",  // 提示信息
    "data": {
        "QA1": {
            "question": "Where did you meet your first girlfriend?",
            "answer": "In Zhuhai"
        },
        "QA2": {
            "question": "What is your favourate color?",
            "answer": "Silver"
        },
        "QA3": {
            "question": "What is the first film you watched in therater?",
            "answer": "Interstellar"
        }
    }  // 返回内容
}
```

​    

### 4.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  





## 5 Change Password Interface 

### 5.1 Interface Description    

- Help user change password
- New password should not be the same as the old one

### 5.2 Address  

`{apiAddress}/api/user/change-password`  

### 5.3 Request Type  

**POST**  

### 5.4 Request Parameters  

#### 5.4.1 Header Parameters  

| Key       | Must | Type/Value      | Note         |
| ------------ | ---- | ---------------- | ------------ |
| Content-Type | Yes   | application/json | Request parameter type |

#### 5.4.2 Body Parameters  

| Key    | Mush | Type   | Limit        | Note     |
| --------- | ---- | ------ | --------------- | -------- |
| oldPassword   | Yes   | string | 1 < length < 50 | the old password |
| newPassword  | Yes   | string | 1 < length < 50 | the new password     |
| userId | Yes   | string | 1 < length < 50     | the userId to locate user   |

**Special Note**: passcode encription with SHA256  (* to be decided)

​    

### 5.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "Passord Reset Successful",  // 提示信息
    "data": null  // 返回内容
}
```

​    

### 5.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  



## 6 Forget Password Interface 

### 6.1 Interface Description    

- User find his password by using user questions and account name 

### 6.2 Address  

`{apiAddress}/api/user/forget-password`  

### 6.3 Request Type  

**POST**  

### 6.4 Request Parameters  

#### 6.4.1 Header Parameters  

| Key       | Must | Type/Value      | Note         |
| ------------ | ---- | ---------------- | ------------ |
| Content-Type | Yes   | application/json | Request parameter type |

#### 6.4.2 Body Parameters  

| Key    | Mush | Type   | Limit        | Note     |
| --------- | ---- | ------ | --------------- | -------- |
| account   | No  | string | 1 < length < 50 | Account |
| email   | No  | string | 1 < length < 50 | email binded with the account |
| cellphone   | No  | string | 1 < length < 50 | phone# binded with the account |
| QAPairs  | Yes  | object | 1 < length < 50 | Passord     |
| checkCode | Yes  | string | length = 6      | Check Code   |

**Special Note**:
1. passcode encription with SHA256  (* to be decided)
2. account, email and phoneNum cannot be None at the same time
3. Verification of the triple None logic should be put at frontend
4. QAPairs : 
```json
{
    "QA1": {
            "question": "sample question 1",
            "answer": "sample answer1"
        },
    "QA2": {
            "question": "sample question 2",
            "answer": "sample answer2"
        },
    "QA3": {
            "question": "sample question 3",
            "answer": "sample answer3"
        }
}
```
**Other Interfaces needs to be Involved**:  

| Interface Name   | Interface Address                               | Purpose           |
| ---------- | -------------------------------------- | ------------------ |
| Get Checkcode | `{apiAddress}/api/common/getCheckcode` | getNecessary Checkcode |

​    

### 6.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "User Identity Verified",  // 提示信息
    "data": {
        "user_id": "32d23ab3b420ef"
    }  // 返回内容
}
```

​    

### 6.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  