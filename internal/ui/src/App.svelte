<script lang="ts">
    import {getProjects} from "./api/project";

    let projects: any = $state([])
    async function loadProjects() {
        try {
            projects = await getProjects()
        } catch(e) {
            console.log(e)
        }
    }

    getProjects()
</script>

{#each projects as project (project)}
    { project.name }
{/each}

