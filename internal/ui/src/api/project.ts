import {API_URL} from "./api";

export async function getProjects() {
    const res = await fetch(`${API_URL}/api/project`, {
        method: "GET"
    })

    if (!res.ok) {
        throw new Error(`Failed retrieving projects: ${res.status}`)
    }
}
