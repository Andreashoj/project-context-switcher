<script lang="ts">
    let projects = $state([])
    async function getProjects() {
        const res = await fetch("http://localhost:8080/api/projects", {
            method: "GET"
        })

        if (res.ok) {
            projects = await res.json()
        }
    }

    getProjects()
</script>

{ projects }
