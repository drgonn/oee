export interface TableListItem {
  id?: number;
  uid: string;
  username: string;
  phone: string;
  email: string;
  emailbind: boolean;
  company: string;
  address: string;
  url: string;
  nickname: string;
  headimgurl: string;
  createDate: Date;
}

export interface TablePutItem {
  id?: number;
  username?: string;
  phone?: string;
  email?: string;
  emailbind?: boolean;
  company?: string;
  address?: string;
  url?: string;
  nickname?: string;
  headimgurl?: string;
}

export interface TableListPagination {
  total: number;
  pageSize: number;
  current: number;
}

export interface TableListData {
  list: TableListItem[];
  pagination: Partial<TableListPagination>;
}

export interface TableListParams {
  status?: string;
  name?: string;
  desc?: string;
  key?: number;
  pageSize?: number;
  currentPage?: number;
  filter?: { [key: string]: any[] };
  sorter?: { [key: string]: any };
}
