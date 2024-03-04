<script>
    // import axios from "$lib/axios";
    import axios from "axios";
    import { onMount } from "svelte";
    import Cookies from "js-cookie";
    let data, failure, criteria, order, category, keyword,  allCat, isAdmin;
    onMount(async () => {
        try{
            // get all categories
            const resp = await axios.get('http://localhost:8000/api/categories');
            allCat = resp.data;
            criteria = "eventDate"
            order = "ASC";
            keyword = "";
            category = allCat[0]["Id"]
            // console.log(allCat[0]["Id"]);

            // get all events
            const res = await axios.get(`http://localhost:8000/api/events?criteria=${criteria}&order=${order}&category=${category}&keyword=${keyword}`)
            // const res = await axios.get('http://localhost:8000/api/events');
            data = res.data;
            console.log(res.data);


            if(Cookies.get('token') != undefined){
                const config = {
                    headers: { Authorization: `Bearer ${Cookies.get('token')}` }
                };
                const temp = await axios.post('/admin', {}, config);
                console.log(temp.data);
                isAdmin = temp.data;
            }

            // const res = await fetch('http://localhost:8000/api/events');
            // data = await res.json();
            // console.log(data);
        }catch (e){
            console.error("Error:", e.response);
            // failure = e.response.data.error;
            isAdmin = e.response.data;
        }
    })

    async function refresh(){
        // get all events
        const res = await axios.get(`http://localhost:8000/api/events?criteria=${criteria}&order=${order}&category=${category}&keyword=${keyword}`)
        // const res = await axios.get('http://localhost:8000/api/events');
        data = res.data;
        console.log(res.data);
    }

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

<main class="flex col cen">
    <div class="flex row margin-t cen">
        <!-- Filter:  -->
        <select name="" id="" bind:value={category} on:change={refresh}>
            <option value="">All</option>
            {#if allCat}
                {#each allCat as category}
                    <option value={category.Id}>{category.CategoryName}</option>
                {/each}
            {/if}
        </select>
        <!-- Sort By: -->
        <select name="" id="" bind:value={criteria} on:change={refresh}>
            <option value="eventDate">Event Date</option>
            <option value="eventName">Event Name</option>
            <option value="eventCapacity">Event Capacity</option>
        </select>
        <select name="" id="" bind:value={order} on:change={refresh}>
            <option value="ASC">Ascending</option>
            <option value="DESC">Descending</option>
        </select>
        <input type="text" name="" id="" bind:value={keyword} on:change={refresh} placeholder="Keyword">
    </div>
    {#if data != "" && data != undefined}
    <h1 class="title">Newest Events:</h1>
        {#each data as events}
        <a href="/event/{events.Id}">
            <div class="flex row card border">
                <img src={events.EventImage} alt="" loading="lazy">
                <div class="flex col desc">
                    <h1>{events.EventName}</h1>
                    <div class="flex flex-end">
                        <p>{events.EventDescription}</p>
                        <div class="flex row left">
                            <p class="icon gap flex"><iconify-icon icon="bi:geo-alt-fill"></iconify-icon>{events.EventLocation}</p>
                            <p class="icon gap flex"><iconify-icon icon="bi:calendar-week"></iconify-icon>{events.EventDate}</p>
                        </div>
                        {#if isAdmin} 
                        <div class="flex row test">
                            <a href={"/event/"+events.Id+"/edit"} class="flex cen button"><iconify-icon icon="bi:pencil"></iconify-icon></a>
                            <button on:click={() => deleteEvent(events.Id)} class="button"><iconify-icon icon="bi:trash"></iconify-icon></button>
                        </div>
                        {/if}
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
    /* .test{
        position: relative;
        bottom: -45%;
    } */
    img{
        width: 20rem;
    }
    h1{
        height: 3rem;
    }
    input{
        border: none;
        outline: none;
        border-radius: 0.2rem;
        background-color: var(--fg);
        font-size: 12pt;
        height: 2rem;
    }
    select{
        /* margin-top: 5rem; */
        margin: 0.5rem;
        height: 2rem;
        border-radius: 0.2rem;
        background-color: var(--fg);
        border: none;
        outline: none;
    }
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
    .gap{
        gap: 1rem;
    }
    .icon{
        align-items: center;
    }
    .left{
        justify-content: center;
        align-items: flex-start;
    }
    .desc{
        margin-left: 1rem;
        width: 20rem;
        justify-content: center;
    }
    .card{
        /* margin-left: 4rem; */
        /* gap: 1rem; */
        margin-bottom: 1rem;
        margin-top: 1rem;
        background-color: var(--fg);
        width: 62rem;
        border-radius: 0.4rem;
    }
    .actions, .title{
        margin-left: 4rem;
    }
    .actions>a{
        margin: 0.5rem 0.7rem 0.5rem 0.7rem;
        background-color: var(--fg);
        padding: 0.7rem;
        border-radius: 0.4rem;
    }
    div>p {
        margin-top: 1rem;
        width: 37rem;
    }
    button{
        /* margin-top: 3rem;
        margin-left: 1rem; */
        outline: none;
        border: none;
        padding: 0.75rem;
        background-color: var(--button-fill);
        color: white;
        border-radius: 0.4rem;
        margin: 0.25rem;
    }
    .margin-t{
        margin-top: 5rem;
    }
</style>