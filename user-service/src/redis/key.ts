export const limitKey = 60;
export const usersKey = (userId: string) => 'users#' + userId;
export const allUserKey = (page: number, limit: number, sort: string) =>
  `all_users:${page}:${limit}:${sort}`;
