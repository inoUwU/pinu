# 現在のアーキテクチャ図

**作成日**: 2025年7月29日  
**プロジェクト**: Pinu  
**アーキテクチャパターン**: ポート&アダプター（ヘキサゴナルアーキテクチャ）

## システム全体のアーキテクチャ図

```mermaid
graph TB
    %% External Systems
    Client[Web Browser<br/>Frontend]
    DB[(Database<br/>SQLite)]
    
    %% Application Layers
    subgraph "Application (Backend)"
        subgraph "Presentation Layer"
            Controller[Controllers<br/>user_controller.go]
            Routes[API Routes<br/>Fiber Router]
        end
        
        subgraph "Application Layer"
            Usecase[UseCases<br/>user_usecase.go<br/>+ Validation<br/>+ Business Logic]
            InputDTO[Input DTOs<br/>get_users_input.go]
            OutputDTO[Output DTOs<br/>get_users_output.go]
        end
        
        subgraph "Domain Layer"
            Entity[Entities<br/>user.go]
            RepoInterface[Repository Interface<br/>IUserRepository<br/>【ポート】]
        end
        
        subgraph "Infrastructure Layer"
            RepoImpl[Repository Implementation<br/>user_repository_impl.go<br/>【アダプター】]
            DI[Dependency Injection<br/>injection.go]
        end
    end
    
    %% Data Flow
    Client -->|HTTP Request| Routes
    Routes --> Controller
    Controller --> Usecase
    Usecase --> InputDTO
    Usecase --> RepoInterface
    RepoInterface <-.- RepoImpl
    RepoImpl --> DB
    
    %% Response Flow
    DB --> RepoImpl
    RepoImpl --> RepoInterface
    RepoInterface --> Usecase
    Usecase --> OutputDTO
    OutputDTO --> Controller
    Controller --> Routes
    Routes -->|HTTP Response| Client
    
    %% DI Container
    DI -.-> Controller
    DI -.-> Usecase
    DI -.-> RepoImpl
    
    %% Entity Usage
    Entity -.-> Usecase
    Entity -.-> RepoInterface
    Entity -.-> OutputDTO
    
    %% Styling
    classDef primaryAdapter fill:#e1f5fe,stroke:#01579b,stroke-width:2px
    classDef application fill:#f3e5f5,stroke:#4a148c,stroke-width:2px
    classDef domain fill:#e8f5e8,stroke:#1b5e20,stroke-width:2px
    classDef secondaryAdapter fill:#fff3e0,stroke:#e65100,stroke-width:2px
    classDef external fill:#ffebee,stroke:#c62828,stroke-width:2px
    
    class Controller,Routes primaryAdapter
    class Usecase,InputDTO,OutputDTO application
    class Entity,RepoInterface domain
    class RepoImpl,DI secondaryAdapter
    class Client,DB external
```

## レイヤー別の責任詳細図

```mermaid
graph LR
    subgraph "Controllers (Primary Adapter)"
        C1[HTTP Request Handling]
        C2[Response Formatting]
        C3[Validation of Input Format]
        C4[Error Handling]
    end
    
    subgraph "UseCases (Application Layer)"
        U1[Business Logic]
        U2[Data Validation]
        U3[Orchestration]
        U4[Transaction Management]
        U5[Input/Output Transformation]
    end
    
    subgraph "Repository Interface (Domain Port)"
        R1[Data Access Contract]
        R2[Domain Model Definition]
        R3[Query Interface]
    end
    
    subgraph "Repository Implementation (Secondary Adapter)"
        I1[SQL Query Execution]
        I2[Database Connection]
        I3[ORM Operations]
        I4[Data Mapping]
    end
    
    C1 --> U1
    C2 --> U5
    C3 --> U2
    C4 --> U4
    
    U1 --> R1
    U2 --> R2
    U3 --> R3
    
    R1 --> I1
    R2 --> I4
    R3 --> I2
    
    classDef controller fill:#e1f5fe,stroke:#01579b,stroke-width:2px
    classDef usecase fill:#f3e5f5,stroke:#4a148c,stroke-width:2px
    classDef port fill:#e8f5e8,stroke:#1b5e20,stroke-width:2px
    classDef adapter fill:#fff3e0,stroke:#e65100,stroke-width:2px
    
    class C1,C2,C3,C4 controller
    class U1,U2,U3,U4,U5 usecase
    class R1,R2,R3 port
    class I1,I2,I3,I4 adapter
```

## 依存性の方向図

```mermaid
graph TB
    subgraph "外部システム"
        Frontend[Frontend<br/>Next.js]
        Database[(Database<br/>SQLite)]
    end
    
    subgraph "アプリケーション内部"
        Controller[Controller]
        Usecase[UseCase]
        RepoInterface[Repository<br/>Interface]
        RepoImpl[Repository<br/>Implementation]
    end
    
    %% Dependencies (arrows show dependency direction)
    Frontend --> Controller
    Controller --> Usecase
    Usecase --> RepoInterface
    RepoImpl --> RepoInterface
    RepoImpl --> Database
    
    %% Invocation (dotted lines show call direction)
    Controller -.-> Usecase
    Usecase -.-> RepoInterface
    RepoInterface -.-> RepoImpl
    
    classDef external fill:#ffebee,stroke:#c62828,stroke-width:2px
    classDef internal fill:#e8f5e8,stroke:#1b5e20,stroke-width:2px
    
    class Frontend,Database external
    class Controller,Usecase,RepoInterface,RepoImpl internal
```

## ファイル構造とマッピング

```mermaid
graph TB
    subgraph "Project Structure"
        subgraph "Controllers"
            UC[user_controller.go]
        end
        
        subgraph "UseCases"
            UUC[user_usecase.go]
            Input[input/<br/>get_users_input.go]
            Output[output/<br/>get_users_output.go]
        end
        
        subgraph "Domain"
            Entity[entities/<br/>user.go]
            RepoInt[repositories/<br/>user_repository.go]
        end
        
        subgraph "Infrastructure"
            RepoImpl[repositories/<br/>user_repository_impl.go]
            Injection[middleware/<br/>injection.go]
        end
    end
    
    UC --> UUC
    UUC --> Input
    UUC --> Output
    UUC --> RepoInt
    UUC --> Entity
    RepoImpl --> RepoInt
    RepoImpl --> Entity
    
    Injection -.-> UC
    Injection -.-> UUC
    Injection -.-> RepoImpl
    
    classDef controller fill:#e1f5fe,stroke:#01579b,stroke-width:2px
    classDef usecase fill:#f3e5f5,stroke:#4a148c,stroke-width:2px
    classDef domain fill:#e8f5e8,stroke:#1b5e20,stroke-width:2px
    classDef infrastructure fill:#fff3e0,stroke:#e65100,stroke-width:2px
    
    class UC controller
    class UUC,Input,Output usecase
    class Entity,RepoInt domain
    class RepoImpl,Injection infrastructure
```

## アーキテクチャの特徴

### ✅ 採用した原則

1. **依存性逆転の原則**: 内側の層は外側の層に依存しない
2. **単一責任の原則**: 各層は明確な責任を持つ
3. **インターフェース分離の原則**: 適切なポートを定義
4. **開放閉鎖の原則**: 実装を変更せずに拡張可能

### 🔧 技術的な決定

1. **ポインタ型の使用**: サービス・リポジトリはポインタ型で統一
2. **インターフェース直接使用**: `*Interface` ではなく `Interface` 型を使用
3. **サービス層の削除**: 不要な抽象化を排除
4. **バリデーションの配置**: ユースケース層でビジネスルールを実装

### 📈 利点

- **テスタビリティ**: 各層を独立してテスト可能
- **保守性**: 責任が明確で変更の影響範囲が限定的
- **拡張性**: 新機能追加時の一貫したパターン
- **理解しやすさ**: シンプルで直感的な構造

この設計により、スケーラブルで保守性の高いGo言語アプリケーションが実現されています。
