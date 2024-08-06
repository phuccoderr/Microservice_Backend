export const customerKey = (id: string) => `customers#${id}`
export const allCustomerKey = (page: number, limit: number, sort: string) =>
  `all_customers:${page}:${limit}:${sort}`;