const prefix = "";

export const USERS = `${prefix}/users/v1`;
export const USERS_ME = `${prefix}/users/v1/me`;
export const USERS_LOGIN = `${prefix}/users/v1/me/login`;
export const USERS_LOGOUT = `${prefix}/users/v1/me/logout`;

export const CONFIGURATIONS = `${prefix}/configurations`;
export const CONFIGURATIONS_ADD = `${CONFIGURATIONS}/v1`;
export const CONFIGURATIONS_UPDATE = `${CONFIGURATIONS}/v1/:id`;
export const CONFIGURATIONS_DELETE = `${CONFIGURATIONS}/v1/:id`;
export const CONFIGURATIONS_LIST = `${CONFIGURATIONS}/v1`;
export const CONFIGURATIONS_LIST_AVAILABLE = `${CONFIGURATIONS}/v1/available`;

export const ROUTERS = `${prefix}/routers`;

export const RANDOM_KEYS = `${prefix}/random-keys`;
