export const allUserKey = (page: number, limit: number, sort: string) =>
  `all_users:${page}:${limit}:${sort}`;
