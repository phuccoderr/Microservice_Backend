export class ResponsePaginationDTO {
  total_items: number;
  total_pages: number;
  current_page: number;
  start_count: number;
  end_count: number;
  entities: any[];
}