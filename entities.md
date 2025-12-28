# Mapeamento de Entidades - Solis Hub

Este documento descreve todas as entidades do sistema Solis Hub com seus respectivos atributos e relacionamentos.

## 1. KPI (Indicadores-Chave de Performance)

### KpiCategory
```typescript
interface KpiCategory {
  id: string;
  title: string;
  data: MonthlyData[];
  color: string;
  unit?: 'currency' | 'percent' | 'number';
  logs?: KpiLog[];
}
```

### MonthlyData
```typescript
interface MonthlyData {
  month: string;
  realized: number | null;
  meta: number | null;
  breakdown?: BreakdownItem[];
}
```

### BreakdownItem
```typescript
interface BreakdownItem {
  label: string;
  value: number;
  subItems?: BreakdownItem[];
}
```

### KpiLog
```typescript
interface KpiLog {
  id: string;
  date: string;       // ISO string
  timestamp: string;  // Hora formatada
  user: string;
  month: string;
  oldValue: number | null;
  newValue: number;
  action: 'update' | 'create';
  context?: string;
}
```

## 2. Tarefas (Task Management)

### Task
```typescript
interface Task {
  id: string;
  title: string;
  description?: string;
  startDate?: string;
  dueDate: string; // 'Hoje', 'Amanhã', ou data específica
  priority: 'alta' | 'media' | 'baixa';
  status: 'todo' | 'in_progress' | 'done';
  category: 'comercial' | 'marketing' | 'admin';
  assignee?: 'Jackson' | 'Beatriz' | 'Larissa' | string;
  flows?: string[];
  subtasks?: Subtask[];
  comments?: Comment[];
  archived?: boolean;
}
```

### Subtask
```typescript
interface Subtask {
  id: string;
  title: string;
  completed: boolean;
  assignee?: string;
  dueDate?: string;
}
```

### Comment
```typescript
interface Comment {
  id: string;
  user: string;
  text: string;
  timestamp: string;
}
```

## 3. Notificações

### Notification
```typescript
interface Notification {
  id: string;
  type: 'mention' | 'deadline' | 'system';
  title: string;
  message: string;
  timestamp: string;
  read: boolean;
  taskId?: string;
  archived?: boolean;
}
```

## 4. Representantes Comerciais

### RepresentativeProfile
```typescript
interface RepresentativeProfile {
  code: number;
  company: string;
  name: string; // Nome Completo / Display Name
  region: string;
  phone: string;
  email: string;
  attendant: string;
  stats?: {
    trainingCount: number;
    onlineCount: number;
    offlineCount: number;
    offlineValue: number;
  };
}
```

### RepData
```typescript
interface RepData {
  name: string;
  [month: string]: number | string; // Chaves de mês com valores
}
```

### RepTableData
```typescript
interface RepTableData {
  title: string;
  columns: string[];
  rows: {
    month: string;
    target: number;
    values: Record<string, number>;
  }[];
}
```

## 5. Marketing Digital

### ChannelPerformance
```typescript
interface ChannelPerformance {
  channel: string;
  visits: number;
  leads: number;
  conversion: number; // percentage
  investment: number;
  cpl: number;
  roas: number;
}
```

### MarketingChannelData
```typescript
interface MarketingChannelData {
  [month: string]: ChannelPerformance[];
}
```

## 6. PDV (Ponto de Venda)

### PdvPost
```typescript
interface PdvPost {
  id: string;
  repName: string;
  pdvName: string;
  postDate: string; // YYYY-MM-DD
  month: string;
  platform: 'instagram' | 'facebook' | 'linkedin' | 'tiktok';
  link: string;
  proofUrl?: string; // URL da imagem ou anexo
  status: 'verified' | 'pending';
}
```

### RecurrentPdv
```typescript
interface RecurrentPdv {
  id: string;
  name: string;
  repName: string;
  city?: string;
  followers?: number;
  instagramProfile?: string;
}
```

### ShowroomItem
```typescript
interface ShowroomItem {
  id: string;
  pdv: string;
  city: string;
  contact: string;
  repName: string;
  notDelivered: boolean; // Coluna "Não entregue"
  deliveryForecast: string; // "Previsão de entrega"
  delivered: boolean; // Coluna "Já entregue"
  workshopDate: string; // "Data do Workshop"
}
```

## 7. Ações Offline

### OfflineAction
```typescript
interface OfflineAction {
  solicitado: string;
  data: string;
  aprovado: string;
  pedido: string;
  saida: string;
  previsao: string;
  entrega: string;
  cidade: string;
  uf: string;
  pontuado: string;
  status: string;
  pdv?: string; // Nome da Loja
  responsavel?: string; // Representante responsável
  // Campos adicionais para formulário detalhado
  contact?: string;
  phone?: string;
  email?: string;
  proposal?: string;
  purchaseVolume?: string;
  category?: string; // Categoria da ação (ex: Brindes, Propaganda)
}
```

## 8. Gestão de Brindes

### GiftItem
```typescript
interface GiftItem {
  name: string;
  stock: number;
  price: number;
}
```

### GiftTransaction
```typescript
interface GiftTransaction {
  id: string;
  itemName: string;
  date: string;
  time: string;
  quantity: number;
  unit: string;
  type: 'in' | 'out'; // in = Entrada, out = Saída
  price?: number; // Apenas para entradas
  representative?: string; // Apenas para saídas
}
```

## 9. Financeiro

### AccountPayable
```typescript
interface AccountPayable {
  id: string;
  supplier: string;
  dueDate: string; // YYYY-MM-DD
  description: string;
  amount: number;
  nfArrived: boolean;
  boletoArrived: boolean;
  status: 'pending' | 'sent_to_finance';
  recurrence?: 'monthly' | 'yearly' | 'none';
}
```

## 10. Calendário de Conteúdo

### CalendarPost
```typescript
interface CalendarPost {
  id: string;
  date: string; // YYYY-MM-DD
  time: string; // HH:MM
  title: string;
  description?: string; 
  caption: string;
  category: 'official' | 'solis_voce' | 'leonardo' | 'luiz';
  type: 'video' | 'static' | 'carousel' | 'story' | 'article_linkedin' | 'article_blog';
  status: 'in_progress' | 'review' | 'adjust' | 'approved' | 'published';
  image?: string; 
  platforms?: string[]; 
  publishedPlatforms?: string[]; 
  assignee?: string; 
  history: PostHistoryEvent[];
}
```

### PostHistoryEvent
```typescript
interface PostHistoryEvent {
  id: string;
  action: 'upload' | 'adjust_request' | 'approved' | 'published' | 'status_change';
  user: string;
  text?: string;
  timestamp: string;
}
```

## 11. Social Media Benchmarking

### SocialBenchmarking
```typescript
interface SocialBenchmarking {
  brand: string;
  avgLikes: number;
  avgComments: number;
  followers?: number;
}
```

## 12. Credenciais e Contatos

### ProgramCredential
```typescript
interface ProgramCredential {
  name: string;
  user?: string;
  password?: string;
  access?: string;
  notes?: string;
}
```

### InternalContact
```typescript
interface InternalContact {
  name: string;
  role: string;
  department: string;
  extension: string; // Usando string pois alguns podem ter caracteres especiais
  email: string;
}
```

## 13. Orçamento

### BudgetItem
```typescript
interface BudgetItem {
  id: string;
  codObj: string;
  obj: string;
  codGrp: string;
  grp: string;
  cod: string;
  desc: string;
  vals: number[]; // Orçado
  realizedVals: number[]; // Realizado
}
```

## Relacionamentos Principais

1. **KpiCategory** → **MonthlyData** (1:N)
2. **MonthlyData** → **BreakdownItem** (1:N)
3. **KpiCategory** → **KpiLog** (1:N)
4. **Task** → **Subtask** (1:N)
5. **Task** → **Comment** (1:N)
6. **Notification** → **Task** (N:1, opcional)
7. **CalendarPost** → **PostHistoryEvent** (1:N)
8. **RepresentativeProfile** → **PdvPost** (1:N)
9. **RepresentativeProfile** → **RecurrentPdv** (1:N)
10. **RepresentativeProfile** → **ShowroomItem** (1:N)
11. **GiftItem** → **GiftTransaction** (1:N)

## Tipos de Dados Comuns

- **Datas**: Geralmente no formato YYYY-MM-DD ou strings ISO
- **IDs**: Strings únicas (geralmente UUIDs)
- **Status**: Enums específicos para cada entidade
- **Moedas**: Numbers (geralmente em reais)
- **Percentagens**: Numbers (0-100)