
// using fetch API and promises
fetch("https://community-boss-api-development.up.railway.app/beta/quarters").
  then(res => res.json()).
  then(parsedResponse => {
    console.log(parsedResponse);
  });

/*
 * Of course, you can also use the "await" keyword instead
 *  of handling resolving the promise with "then" if you're
 *  in an async function or your project supports
 *  asynchronous code top-level in your javascript.
 */
