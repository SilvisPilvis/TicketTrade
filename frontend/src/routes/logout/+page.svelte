<script>
    import { onMount } from 'svelte';
    onMount(async () => {
        console.log("User logged out");
        window.location = "/"
	});
</script>