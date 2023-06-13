<script lang="ts">
  import { CodeBlockLanguage } from "../models/codeBlocks";
  import Highlight from "svelte-highlight";
  import LineNumbers from "svelte-highlight";
  import typescript from "svelte-highlight/languages/typescript";
  import python from "svelte-highlight/languages/python";
  import json from "svelte-highlight/languages/json";
  import a11yDark from "svelte-highlight/styles/a11y-dark";
  // import a11yDark from "svelte-highlight/styles/a11yDark";
  import { githubDark } from "svelte-highlight/styles";

  export let url: string;

  const langMap: any = {
    "js": typescript,
    "ts": typescript,
    "python": python,
  };

  const languages: CodeBlockLanguage[] = [
    new CodeBlockLanguage("js", "const response = await fetch(\"[REQUEST_URI]\");\nconst jsonData = await response.json();\nconsole.log(jsonData);"),
    new CodeBlockLanguage("python", "import requests\n\nresponse = requests.get(\"[REQUEST_URI]\")\nresponse_json = response.json()"),
  ];

  let tab: number = 0;

  let requestOutput: string | null = null;
  let loadingResponse: boolean = false;

  function makeRequest() {
    console.log(url)
    loadingResponse = true;
    fetch("https://" + url, {
      headers: {"Access-Control-Allow-Origin": "*"}
    }).then(response => {
      response.json().then(j => {
        requestOutput = j;
        loadingResponse = false;
      });
    })
  }

</script>

<svelte:head>
  {@html githubDark}
</svelte:head>

<div class="flex flex-row space-x-4 pt-4">
  {#each languages as lang, i}
    <button class="text-white font-bold" on:click={() => tab = i}>{lang.lang.toUpperCase()}</button>
  {/each}
</div>

<div class="bg-[#0D1117] rounded-lg border-white border-2 px-2 py-2">
  <Highlight language={langMap[languages[tab].lang]} code={languages[tab].formattedCode(url)} />
</div>
{#if loadingResponse}
  Loading
  {:else if requestOutput}
    <Highlight language={json} code={requestOutput}>
    </Highlight>
{/if}

<!-- <button on:click={makeRequest}>Test Endpoint</button> -->
