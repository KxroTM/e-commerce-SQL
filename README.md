# Warning !

The script must be executed on an empty database.

# Creation of My Tables

- Users Table

```sql
CREATE TABLE IF NOT EXISTS user (
    client_id INT AUTO_INCREMENT PRIMARY KEY,
    nom VARCHAR(100) NOT NULL,
    prenom VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    telephone VARCHAR(15),
    hashed_password VARCHAR(255) NOT NULL,
    role_id INT,
    photo_id INT,
    FOREIGN KEY (role_id) REFERENCES role(role_id),
    FOREIGN KEY (photo_id) REFERENCES photo(photo_id)
);
```

- Sellers Table

```sql
CREATE TABLE IF NOT EXISTS vendeur (
    vendeur_id INT AUTO_INCREMENT PRIMARY KEY,
    client_id INT,
    date_creation DATETIME DEFAULT CURRENT_TIMESTAMP,
    nom_commercial VARCHAR(255) NOT NULL,
    type_activite VARCHAR(255) NOT NULL,
    n_registre_commerce VARCHAR(255) NOT NULL,
    n_tva VARCHAR(255) NOT NULL,
    telephone VARCHAR(15),
    email VARCHAR(255) NOT NULL,
    FOREIGN KEY (client_id) REFERENCES user(client_id),
);
```

- Addresses Table

```sql
CREATE TABLE IF NOT EXISTS address (
    adresse_id INT AUTO_INCREMENT PRIMARY KEY,
    client_id INT,
    numero_rue INT NOT NULL,
    rue VARCHAR(255) NOT NULL,
    ville VARCHAR(100) NOT NULL,
    code_postal VARCHAR(10) NOT NULL,
    pays VARCHAR(100) NOT NULL,
    region VARCHAR(100) NOT NULL,
    complement_adresse TEXT,
    FOREIGN KEY (client_id) REFERENCES user(client_id)
);
```

- Products Table

```sql
CREATE TABLE IF NOT EXISTS product (
    produit_id INT AUTO_INCREMENT PRIMARY KEY,
    nom VARCHAR(255) NOT NULL,
    prix DECIMAL(10, 2) NOT NULL,
    description TEXT,
    stock INT NOT NULL,
    vendeur_id INT,
    etat_id INT,
    photo_id INT,
    FOREIGN KEY (etat_id) REFERENCES etat_produit(etat_id),
    FOREIGN KEY (photo_id) REFERENCES photo(photo_id),
    FOREIGN KEY (vendeur_id) REFERENCES vendeur(vendeur_id)
);
```

- Cart Table

```sql
CREATE TABLE IF NOT EXISTS cart (
    panier_id INT AUTO_INCREMENT PRIMARY KEY,
    client_id INT,
    montant DECIMAL(10, 2),
    date_creation DATETIME DEFAULT CURRENT_TIMESTAMP,
    status_id INT,
    FOREIGN KEY (client_id) REFERENCES user(client_id)
    FOREIGN KEY (status_id) REFERENCES status_panier(status_id)
);
```

- Link Table Between Cart and Products

```sql
CREATE TABLE IF NOT EXISTS content_product (
    panier_id INT,
    produit_id INT,
    quantite INT,
    PRIMARY KEY (panier_id, produit_id),
    FOREIGN KEY (panier_id) REFERENCES cart(panier_id),
    FOREIGN KEY (produit_id) REFERENCES product(produit_id)
);
```

- Orders Table

```sql
CREATE TABLE IF NOT EXISTS command (
    commande_id INT AUTO_INCREMENT PRIMARY KEY,
    client_id INT,
    montant DECIMAL(10, 2),
    status_id INT,
    date_livraison DATE NOT NULL,
    paiement_id INT,
    FOREIGN KEY (client_id) REFERENCES user(client_id),
    FOREIGN KEY (status_id) REFERENCES commade_status(status_id),
    FOREIGN KEY (paiement_id) REFERENCES payment(paiement_id)
);
```

- Invoices Table

```sql
CREATE TABLE IF NOT EXISTS invoices (
    invoice_id INTEGER PRIMARY KEY AUTOINCREMENT,
    commande_id INTEGER NOT NULL,
    client_id INTEGER NOT NULL,
    date_facture TEXT DEFAULT CURRENT_TIMESTAMP,
    total_amount REAL NOT NULL,
    tax_amount REAL DEFAULT 0.00,
    status_paiment TEXT,
    paiement_id INTEGER,
    FOREIGN KEY (commande_id) REFERENCES command(commande_id),
    FOREIGN KEY (client_id) REFERENCES user(client_id),
    FOREIGN KEY (paiement_id) REFERENCES payment(paiement_id)
);
```

- Product States Table

```sql
CREATE TABLE IF NOT EXISTS etat_produit (
    etat_id INT AUTO_INCREMENT PRIMARY KEY,
    description VARCHAR(255) NOT NULL
);
```

- Cart Status Table

```sql
CREATE TABLE IF NOT EXISTS status_panier (
    status_id INT AUTO_INCREMENT PRIMARY KEY,
    description VARCHAR(255) NOT NULL
);
```

- Order Status Table

```sql
CREATE TABLE IF NOT EXISTS commande_status (
    status_id INT AUTO_INCREMENT PRIMARY KEY,
    description VARCHAR(255) NOT NULL
);
```

- Roles Table

```sql
CREATE TABLE IF NOT EXISTS role (
    role_id INT AUTO_INCREMENT PRIMARY KEY,
    description VARCHAR(255) NOT NULL
);
```

- Payments Table

```sql
CREATE TABLE IF NOT EXISTS payment (
    paiement_id INT AUTO_INCREMENT PRIMARY KEY,
    type VARCHAR(100) NOT NULL
);
```

- Photos Table

```sql
CREATE TABLE IF NOT EXISTS photo (
    photo_id INT AUTO_INCREMENT PRIMARY KEY,
    url VARCHAR(255) NOT NULL
);
```

- Ratings Table (Product Reviews)

```sql
CREATE TABLE IF NOT EXISTS rate (
    evaluation_id INT AUTO_INCREMENT PRIMARY KEY,
    produit_id INT,
    client_id INT,
    note INT CHECK(note >= 1 AND note <= 5),
    commentaire TEXT,
    FOREIGN KEY (produit_id) REFERENCES product(produit_id),
    FOREIGN KEY (client_id) REFERENCES user(client_id)
);
```
