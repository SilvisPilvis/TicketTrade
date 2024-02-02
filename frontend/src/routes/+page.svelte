<script>
    // import axios from "$lib/axios";
    import axios from "axios";
    import { onMount } from "svelte";
    let data, failure;
    onMount(async () => {
        try{
            const res = await axios.get('http://localhost:8000/api/events');
            // const res = await axios.get('/events');
            data = res.data;
            console.log(res.data);

            // const res = await fetch('http://localhost:8000/api/events');
            // data = await res.json();
            // console.log(data);
        }catch (e){
            console.error("Error:", e.response.data.error);
            failure = e.response.data.error;
        }
    })
</script>

<main>
    {#if data != "" && data != undefined}
        {#each data as events}
        <a href="/event/{events.Id}">
            <div>
                <img src={events.EventImage} alt="" loading="lazy">
                <p>{events.EventName}</p>
                <p>{events.EventDescription}</p>
                <p>{events.EventLocation}</p>
                <p>{events.EventDate}</p>
            </div>
        </a>
        {/each}
    {/if}
    {#if failure != "" && failure != undefined}
        <p class="error">{failure}</p>
    {/if}
</main>

<style>
    a{
        text-decoration: none;
        color: black;
    }
</style>