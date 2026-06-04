const BASE_URL = "http://localhost:8080";

const movies = [
  { title: "Inception", producer: "Emma Thomas", director: "Christopher Nolan", release_year: 2010 },
  { title: "The Dark Knight", producer: "Emma Thomas", director: "Christopher Nolan", release_year: 2008 },
  { title: "Interstellar", producer: "Emma Thomas", director: "Christopher Nolan", release_year: 2014 },
  { title: "The Matrix", producer: "Joel Silver", director: "Lana Wachowski", release_year: 1999 },
  { title: "Pulp Fiction", producer: "Lawrence Bender", director: "Quentin Tarantino", release_year: 1994 },
];

const actors = [
  { first_name: "Leonardo", last_name: "DiCaprio", birth_date: "1974-11-11", salary: 20000000 },
  { first_name: "Joseph", last_name: "Gordon-Levitt", birth_date: "1981-02-17", salary: 5000000 },
  { first_name: "Elliot", last_name: "Page", birth_date: "1987-02-21", salary: 4000000 },
  { first_name: "Christian", last_name: "Bale", birth_date: "1974-01-30", salary: 15000000 },
  { first_name: "Heath", last_name: "Ledger", birth_date: "1979-04-04", salary: 10000000 },
  { first_name: "Matthew", last_name: "McConaughey", birth_date: "1969-11-04", salary: 15000000 },
  { first_name: "Anne", last_name: "Hathaway", birth_date: "1982-11-12", salary: 10000000 },
  { first_name: "Keanu", last_name: "Reeves", birth_date: "1964-09-02", salary: 12000000 },
  { first_name: "Carrie-Anne", last_name: "Moss", birth_date: "1967-08-21", salary: 5000000 },
  { first_name: "John", last_name: "Travolta", birth_date: "1954-02-18", salary: 8000000 },
  { first_name: "Uma", last_name: "Thurman", birth_date: "1970-04-29", salary: 7000000 },
];

const budgets = [
  { total_budget: 160000000, fees_in_prod_country: 292576195, fees_in_other: 543000000 },
  { total_budget: 185000000, fees_in_prod_country: 534858444, fees_in_other: 469700000 },
  { total_budget: 165000000, fees_in_prod_country: 188020017, fees_in_other: 487000000 },
  { total_budget: 63000000,  fees_in_prod_country: 171479930, fees_in_other: 292037453 },
  { total_budget: 8000000,   fees_in_prod_country: 107928762, fees_in_other: 106100000 },
];

// Актёры для каждого фильма (по индексу)
const movieActors = [
  [0, 1, 2],    // Inception: DiCaprio, Gordon-Levitt, Page
  [3, 4],       // The Dark Knight: Bale, Ledger
  [5, 6],       // Interstellar: McConaughey, Hathaway
  [7, 8],       // The Matrix: Reeves, Moss
  [9, 10],      // Pulp Fiction: Travolta, Thurman
];

async function getAdminToken() {
  const res = await fetch(`${BASE_URL}/login`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email: "admin_2@kino.local", password: "1234" }),
  });
  const data = await res.json();
  return data.token;
}

async function seed() {
  console.log("🎬 Запускаем сид...");

  const token = await getAdminToken();
  const headers = { "Content-Type": "application/json", Authorization: `Bearer ${token}` };

  // Создаём актёров
  console.log("👤 Создаём актёров...");
  const actorIds = [];
  for (const actor of actors) {
    const res = await fetch(`${BASE_URL}/actors`, { method: "POST", headers, body: JSON.stringify(actor) });
    const data = await res.json();
    actorIds.push(data.id);
    console.log(`  ✓ ${actor.first_name} ${actor.last_name} → id=${data.id}`);
  }

  // Создаём фильмы + бюджет + привязываем актёров
  console.log("🎥 Создаём фильмы...");
  for (let i = 0; i < movies.length; i++) {
    const movieRes = await fetch(`${BASE_URL}/movies`, { method: "POST", headers, body: JSON.stringify(movies[i]) });
    const movie = await movieRes.json();
    console.log(`  ✓ ${movie.title} → id=${movie.id}`);

    // Бюджет
    await fetch(`${BASE_URL}/movies/${movie.id}/budget`, {
      method: "POST", headers, body: JSON.stringify(budgets[i]),
    });
    console.log(`    💰 Бюджет добавлен`);

    // Актёры
    for (const actorIdx of movieActors[i]) {
      await fetch(`${BASE_URL}/movies/${movie.id}/actors`, {
        method: "POST", headers, body: JSON.stringify({ actor_id: actorIds[actorIdx] }),
      });
    }
    console.log(`    🎭 Актёры привязаны: ${movieActors[i].map(idx => actorIds[idx]).join(", ")}`);
  }

  console.log("✅ Готово!");
}

seed().catch(console.error);
