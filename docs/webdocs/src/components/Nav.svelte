<script lang="ts">
  import { onMount } from "svelte";

  let url = '';

  onMount(() => (url = window.location.pathname));
</script>

<div
  class="flex flex-row w-screen justify-between bg-blue-700 py-2 px-2 border-b-2 border-white"
>
  <div class="flex flex-row">
    <a href="/" class="text-white text-2xl font-bold">BOSS API Documentation</a> <p class="text-white">WIP</p>
  </div>
  <select disabled>
    <option value="v1">v1 (beta)</option>
  </select>
</div>
<div class="bg-[#063159] border-b-2 border-red-500">
  <p class="text-white px-2">
    {url}
  </p>
</div>

