import Vue from 'vue'
import axios from 'axios';
import string from "less/lib/less/functions/string";

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

class HttpRequest {

    constructor() {
        this.axios = axios.create();
    }

    async get(path) {
        return this.execute(this.axios.get, path)
    }

    async post(path, message) {
        return this.execute(this.axios.post, path, message, {
            headers: {
                "Content-Type": "application/json"
            }
        })
    }

    async delete(path, message) {
        return this.execute(this.axios.delete, path, message)
    }

    async execute(requestFunction, path, requestMessage, headers) {
        let response = null;
        const responseMessage = new ResponseMessage();

        try {
            response = await requestFunction(
                path,
                typeof requestMessage === 'undefined' ? null : requestMessage.payload,
                headers,
            ).catch((error) => {
                throw (error);
            });

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
