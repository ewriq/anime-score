<script>
  import axios from "axios";
  import { onMount, onDestroy } from "svelte";
  import { writable } from "svelte/store";
  import Toast from "../Toast.svelte";

  let query = "";
  let timer = null;
  let data = [];
  let err = "";
  const api = "http://localhost:3000/api/anime/search";

  function handle() {
    if (timer) {
      clearTimeout(timer);
    }
    timer = setTimeout(doSearch, 300);
  }


  function doSearch() {
    if (query) {
    axios.post(api, { query })
      .then((response) => {
        data = response.data.data;
      })
      .catch((error) => {
        err = error
        showToast(err)
      });
    }
  }

  onMount(() => {
    query = "";
  });

  onDestroy(() => {
    if (timer) {
      clearTimeout(timer);
    }
  });

  const toastVisible = writable(false);
  let toastMessage;

  function showToast(message) {
    toastMessage = message;
    toastVisible.set(true);

    setTimeout(() => {
      toastVisible.set(false);
    }, 3000);
  }

</script>


<main class="">
    <div
      class="w-full mx-auto max-w-screen-xl p-4"
    >
    <form class="max-w-md mx-auto mt-5">
        <label
          for="default-search"
          class="mb-2 text-sm font-medium text-gray-900 sr-only dark:text-white">Anime Ara</label>
        <div class="relative">
          <div class="absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none" >
            <svg class="w-4 h-4 text-gray-500 dark:text-gray-400"
              aria-hidden="true"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 20 20">
              <path stroke="currentColor"
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z"/>
            </svg>
          </div>
          <input bind:value={query}
            type="search"
            id="default-search"
            class="block w-full p-4 ps-10 text-sm text-gray-900 border border-gray-300 rounded-lg bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
            placeholder="Anime Ara"
            required
            on:input={handle}/>
        </div>
      </form>
      {#if query}
      <table
      class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400 mt-10"
    >
      <thead
        class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400"
      >
        <tr>
          <th scope="col" class="px-4 py-2">Adı</th>
          <th scope="col" class="px-4 py-2">Açıklama</th>
        </tr>
      </thead>
      <tbody>
        {#each data as datas}
          <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
            <th
              scope="row"
              class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white"
            >
              <button
                class=" text-blue font-bold py-1 px-2 rounded"
                on:click={() => (window.location.href = `/anime/${datas.name}`)}
              >
                {datas.name}
              </button>
            </th>
            <td class="px-6 py-4">
              {datas.description}
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
      {/if}
   
    </div>
    <Toast message={toastMessage} bind:visible={$toastVisible} />
  </main>
  
