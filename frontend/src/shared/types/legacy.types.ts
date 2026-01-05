
export interface BreakdownItem {
  label: string;
  value: number;
  subItems?: BreakdownItem[];
}

export interface MonthlyData {
  year?: number; // Year (e.g., 2024, 2025)
  month: string;
  realized: number | null;
  meta: number | null;
  breakdown?: BreakdownItem[]; // Added for detailed view (e.g. Channel breakdown)
  logs?: KpiLog[]; // Logs are stored at MonthlyData level, not KPI level
}

export interface KpiLog {
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

export interface KpiCategory {
  id: string;
  title: string;
  slug: string;
  data: MonthlyData[];
  color: string;
  unit?: string; // 'currency' | 'percent' | 'number'
  logs?: KpiLog[];
}

export interface RepData {
  name: string;
  [month: string]: number | string; // Month keys with values
}

export interface RepTableData {
  title: string;
  columns: string[];
  rows: {
    month: string;
    target: number;
    values: Record<string, number>;
  }[];
}

export interface ChannelPerformance {
  channel: string;
  visits: number;
  leads: number;
  conversion: number; // percentage
  investment: number;
  cpl: number;
  roas: number;
}

export interface MarketingChannelData {
  [month: string]: ChannelPerformance[];
}

export type TaskStatus = 'pending' | 'in_progress' | 'completed' | 'cancelled';
export type TaskPriority = 'low' | 'medium' | 'high' | 'urgent';
export type TaskCategory = 'commercial' | 'marketing' | 'training' | 'administrative' | 'support';
export type TaskFlow = 'backlog' | 'to_do' | 'in_progress' | 'review' | 'done';
export type NotificationType = 'mention' | 'deadline' | 'system';
export type AllowedUser = TaskFlow;

export interface Subtask {
  id: string;
  task_id: string;
  title: string;
  description: string;
  priority: TaskPriority;
  status: TaskStatus;
  due_date?: string;
  assignee_id?: string;
  created_at: string;
  updated_at: string;
}

export interface Comment {
  id: string;
  task_id: string;
  user_id: string;
  text: string;
  timestamp: string;
}

export interface Notification {
  id: string;
  user_id: string;
  task_id?: string;
  notification_type: NotificationType;
  title: string;
  message: string;
  read: boolean;
  archived: boolean;
  timestamp: string;
}

export interface Task {
  id: string;
  title: string;
  description: string;
  category: TaskCategory;
  priority: TaskPriority;
  status: TaskStatus;
  flow: TaskFlow;
  due_date?: string;
  assignee_id?: string;
  flows: TaskFlow[];
  archived: boolean;
  subtasks?: Subtask[];
  comments?: Comment[];
  notifications?: Notification[];
  created_at: string;
  updated_at: string;
}

export interface PdvPost {
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

export interface RecurrentPdv {
    id: string;
    name: string;
    repName: string;
    city?: string;
    followers?: number;
    instagramProfile?: string;
}

export interface ShowroomItem {
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

export interface OfflineAction {
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

export interface GiftItem {
    name: string;
    stock: number;
    price: number;
}

export interface GiftTransaction {
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

export interface AccountPayable {
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

export interface RepresentativeProfile {
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

export type PostCategory = 'official' | 'solis_voce' | 'leonardo' | 'luiz';
export type PostType = 'video' | 'static' | 'carousel' | 'story' | 'article_linkedin' | 'article_blog';
export type PostStatus = 'in_progress' | 'review' | 'adjust' | 'approved' | 'published';

export interface PostHistoryEvent {
    id: string;
    action: 'upload' | 'adjust_request' | 'approved' | 'published' | 'status_change';
    user: string;
    text?: string;
    timestamp: string;
}

export interface CalendarPost {
    id: string;
    date: string; // YYYY-MM-DD
    time: string; // HH:MM
    title: string;
    description?: string; 
    caption: string;
    category: PostCategory;
    type: PostType; // Novo campo Tipo
    status: PostStatus;
    image?: string; 
    platforms?: string[]; 
    publishedPlatforms?: string[]; 
    assignee?: string; 
    history: PostHistoryEvent[];
}

export interface SocialBenchmarking {
    brand: string;
    avgLikes: number;
    avgComments: number;
    followers?: number;
}

export interface ProgramCredential {
    name: string;
    user?: string;
    password?: string;
    access?: string;
    notes?: string;
}

export interface InternalContact {
    name: string;
    role: string;
    department: string;
    extension: string; // Using string as some might have special chars or multiple
    email: string;
}

export interface BudgetItem {
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
