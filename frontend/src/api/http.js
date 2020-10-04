import Vue from 'vue'
import axios from 'axios';
import string from "less/lib/less/functions/string";
import store from '@/store'
import _ from 'lodash'

const StatusForbidden = 403;

class ResponseMessage {
    status;
    message;
    data;
    errors;

    constructor(status = false, message = null, data = null, errors = []) {
        this.status = status;
        this.message = message;
        this.data = data;
        this.errors = errors;
    }
}

export class RequestMessage {
    payload;

    constructor(payload) {
        this.payload = payload;
    }
}

const METHOD_POST   = 'POST';
const METHOD_GET    = 'GET';
const METHOD_DELETE = 'DELETE';

class HttpRequest {

    constructor() {
        this.axios = axios.create();
    }

    async get(path) {
        return this.execute(METHOD_GET, path)
    }

    async post(path, message) {
        return this.execute(METHOD_POST, path, message)
    }

    async delete(path, message) {
        return this.execute(METHOD_DELETE, path, message)
    }

    async execute(method, path, requestMessage, headers) {
        console.log(`start execute request`, {requestFunction: method, path, requestMessage, headers});

        const defaultHeaders = {headers: {"Content-Type": "application/json"}};
        if (null !== store.getters['user/getToken']) {
            defaultHeaders.headers["Authorization"] = `Bearer ${store.getters['user/getToken']}`
        }

        let response = null;
        const responseMessage = new ResponseMessage();

        try {
            switch (method) {
                case METHOD_GET:
                    response = await this.axios.get(path, _.merge(defaultHeaders, headers));
                    break;

                case METHOD_POST:
                    response = await this.axios.post(path, requestMessage.payload, _.merge(defaultHeaders, headers));
                    break;

                case METHOD_DELETE:
                    headers = _.merge(defaultHeaders, headers);
                    headers.data = requestMessage.payload;
                    response = await this.axios.delete(path, headers);
                    break;
            }
        } catch (e) {
            response = e.response;

            console.log({
                exception: e,
                response: e.response
            });
        }

        const {data: {data, status, message, errors}} = response;

        responseMessage.data = data || null;
        responseMessage.status = status || false;
        responseMessage.message = message || null;
        responseMessage.errors = errors || [];

        if (!responseMessage.status) {
            if (typeof responseMessage.message === 'string') {
                Vue.$toast.error(responseMessage.message);
            }

            responseMessage.errors.forEach(error => Vue.$toast.error(error))
        }

        console.log({
            path: path,
            requestMessage: requestMessage,
            response: response,
            responseMessage: responseMessage
        });

        return {response, responseMessage}
    }
}

export const httpRequest = new HttpRequest();
