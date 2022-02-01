# ERD
### user domains
```mermaid
erDiagram
    USER_OAUTH_PROVIDERS {
        string id
        string user_id
        enum provider
        time expiry
        string access_token
        string refresh_token
   }
    USERS ||--o{ USER_OAUTH_PROVIDERS : USERS_id
    USERS {
        string id
        enum status
        enum type
        bool signup
        string username
    }
    USERS ||--o{ USER_META : USERS_id
    USER_META {
        string id
        string user_id
        string profile
    }
    USERS ||--o{ USER_DEVICES : USERS_id
    USER_DEVICES {
        string id
        string user_id
        enum status
        enum type
        string os
        string platform
    }
    
    SESSIONS {
        string key
        string user_data
    }
```

### profiles domains
```mermaid
erDiagram
    PROFILES }|--|| USERS : user_id
    PROFILES {
        string id
        string nickname
        string primary
        string image_id
        string description
    }
    PROFILES ||--o{ PROFILE_ACTIONS : profile_id
    PROFILE_ACTIONS {
        string id PK
        string profile_id FK
        string type
        string priority
        string size
    }
    ACTION_ANCHORS {
        string id
        string profile_action_id
        string type
        string url
    }
```
