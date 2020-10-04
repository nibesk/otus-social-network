import consts from './consts'

export default consts

export const splitNumbers = (number) => {
    const parts = number.toString().split('.');
    parts[0] = parts[0].replace(/\B(?=(\d{3})+(?!\d))/g, ' ');

    return parts.join('.');
};

export const localStorageSet = (key, value) => {
    window.localStorage.setItem(key, value)
};

export const localStorageGet = (key) => {
    return window.localStorage.getItem(key)
};

export const localStorageDelete = (key) => {
    window.localStorage.removeItem(key)
};
