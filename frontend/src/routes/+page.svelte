<script>
    // import axios from "$lib/axios";
    import axios from "axios";
    import { onMount } from "svelte";
    let data, failure, criteria, order, category;
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

    function deleteEvent(id){
    axios.delete(`http://localhost:8000/api/event/${id}`)
        .then(function (response) {
            console.log(response);
            success = response.data
        })
        .catch(function (error) {
            console.log(error);
            failed = error.response.data.error;
        });
}
</script>

<main>
    <select name="" id="" bind:value={criteria}>
        <option value="createdAt">Date</option>
        
    </select>
    <select name="" id="" bind:value={category}>
        <option value="eventCategory">Genre</option>
    </select>
    <select name="" id="" bind:value={order}>
        <option value="ASC">Ascending</option>
        <option value="DESC">Descending</option>
    </select>
    {#if data != "" && data != undefined}
    <h1 class="title">Newest Events:</h1>
        {#each data as events}
        <div class="actions flex row">
            <a href={"/event/"+events.Id+"/edit"} class="flex cen">Edit Event</a>
            <button on:click={() => deleteEvent(events.Id)}>Delete Event</button>
        </div>
        <a href="/event/{events.Id}">
            <div class="flex row card">
                <img src={events.EventImage} alt="" loading="lazy">
                <div class="flex col desc">
                    <h1>{events.EventName}</h1>
                    <div class="flex col left">
                        <p>{events.EventDescription}</p>
                        <p>{events.EventLocation}</p>
                        <p>{events.EventDate}</p>
                    </div>
                </div>
            </div>
        </a>
        {/each}
    {/if}
    {#if failure != "" && failure != undefined}
        <p class="error">{failure}</p>
    {/if}
</main>

<style>
    main{
        padding-bottom: 5rem;
    }
    a{
        text-decoration: none;
        color: var(--fg);
    }
    img{
        border-radius: 0.4rem;
        max-height: 30rem;
    }
    .left{
        justify-content: center;
        align-items: flex-start;
    }
    .desc{
        margin-left: 1rem;
        width: 20rem;
    }
    .card{
        margin-left: 4rem;
        margin-bottom: 1rem;
        margin-top: 1rem;
        background-color: var(--bg);
        width: 62rem;
        border-radius: 0.4rem;
    }
    .actions, .title{
        margin-left: 4rem;
    }
    .actions>a, .actions>button{
        margin: 0.5rem 0.7rem 0.5rem 0.7rem;
    }
    .actions>a{
        background-color: var(--bg);
        padding: 0.7rem;
        border-radius: 0.4rem;
    }
    div>p {
        margin-top: 1rem;
        width: 37rem;
    }
    button{
        margin-top: 3rem;
        margin-left: 1rem;
        outline: none;
        border: none;
        padding: 0.75rem;
        background-color: var(--bg);
        color: var(--fg);
        border-radius: 0.4rem;
    }
</style>