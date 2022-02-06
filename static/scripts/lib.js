const BASE_URL = "http://121.4.43.217:8080/";

export class Alexios {
    /**
     * @param {string} path
     * @returns {Promise<string>}
     */
    static async get(path) {
        return (await fetch(`${BASE_URL}api${path}`)).text();
    }
}