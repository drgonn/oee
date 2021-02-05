export interface TableListItem {
  id?: number;
  code: string;
  mean: string;
  cause: string;
  solution: string;
}

export interface TablePutItem {
  id?: number;
  code?: string;
  mean?: string;
  cause?: string;
  solution?: string;
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
