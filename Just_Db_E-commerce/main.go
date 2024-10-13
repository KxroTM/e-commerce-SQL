package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func main() {
	CreateTables()
	LoadFixtures()
	LoadData()
}

func init() {
	var err error
	Db, err = sql.Open("sqlite3", "./Just_Db_E-commerce/e-commerce.sql")
	if err != nil {
		log.Fatal("Erreur lors de l'ouverture de la base de données:", err)
		return
	}
}

func HashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

func CreateTables() {
	query := `
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
        FOREIGN KEY (client_id) REFERENCES user(client_id)
    );

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

    CREATE TABLE IF NOT EXISTS cart (
        panier_id INT AUTO_INCREMENT PRIMARY KEY,
        client_id INT,
        montant DECIMAL(10, 2),
        date_creation DATETIME DEFAULT CURRENT_TIMESTAMP,
        status_id INT,
        FOREIGN KEY (client_id) REFERENCES user(client_id)
		FOREIGN KEY (status_id) REFERENCES status_panier(status_id)
    );

    CREATE TABLE IF NOT EXISTS panier_produit (
        panier_id INT,
        produit_id INT,
        quantite INT,
        PRIMARY KEY (panier_id, produit_id),
        FOREIGN KEY (panier_id) REFERENCES cart(panier_id),
        FOREIGN KEY (produit_id) REFERENCES product(produit_id)
    );

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

    CREATE TABLE IF NOT EXISTS commande_produit (
        commande_id INT,
        produit_id INT,
        quantite INT NOT NULL,
        PRIMARY KEY (commande_id, produit_id),
        FOREIGN KEY (commande_id) REFERENCES command(commande_id),
        FOREIGN KEY (produit_id) REFERENCES product(produit_id)
    );

    CREATE TABLE IF NOT EXISTS etat_produit (
        etat_id INT AUTO_INCREMENT PRIMARY KEY,
        description VARCHAR(255) NOT NULL
    );

	CREATE TABLE IF NOT EXISTS status_panier (
		status_id INT AUTO_INCREMENT PRIMARY KEY,
		description VARCHAR(255) NOT NULL
	);	

    CREATE TABLE IF NOT EXISTS commande_status (
        status_id INT AUTO_INCREMENT PRIMARY KEY,
        description VARCHAR(255) NOT NULL
    );

    CREATE TABLE IF NOT EXISTS role (
        role_id INT AUTO_INCREMENT PRIMARY KEY,
        description VARCHAR(255) NOT NULL
    );

    CREATE TABLE IF NOT EXISTS payment (
        paiement_id INT AUTO_INCREMENT PRIMARY KEY,
        type VARCHAR(100) NOT NULL
    );

    CREATE TABLE IF NOT EXISTS photo (
        photo_id INT AUTO_INCREMENT PRIMARY KEY,
        url VARCHAR(255) NOT NULL
    );

    CREATE TABLE IF NOT EXISTS rate (
        evaluation_id INT AUTO_INCREMENT PRIMARY KEY,
        produit_id INT,
        client_id INT,
        note INT CHECK(note >= 1 AND note <= 5),
        commentaire TEXT,
        FOREIGN KEY (produit_id) REFERENCES product(produit_id),
        FOREIGN KEY (client_id) REFERENCES user(client_id)
    );
    `
	_, err := Db.Exec(query)
	if err != nil {
		log.Fatalf("Erreur lors de la création des tables: %v", err)
	}

	log.Println("Tables créées avec succès")
}

func LoadFixtures() {
	query := `
    INSERT INTO role (role_id, description) VALUES (1, 'admin');
    INSERT INTO role (role_id, description) VALUES (2, 'client');
    INSERT INTO role (role_id, description) VALUES (3, 'vendeur');

    INSERT INTO etat_produit (etat_id, description) VALUES (1, 'Neuf');
    INSERT INTO etat_produit (etat_id, description) VALUES (2, 'Bon état');
    INSERT INTO etat_produit (etat_id, description) VALUES (3, 'Usé');
    INSERT INTO etat_produit (etat_id, description) VALUES (4, 'Reconditionné');

    INSERT INTO status_panier (status_id, description) VALUES (1, 'En cours');
    INSERT INTO status_panier (status_id, description) VALUES (2, 'Validé');
    INSERT INTO status_panier (status_id, description) VALUES (3, 'Annulé');

    INSERT INTO commande_status (status_id, description) VALUES (1, 'En cours de traitement');
    INSERT INTO commande_status (status_id, description) VALUES (2, 'En cours de livraison');
    INSERT INTO commande_status (status_id, description) VALUES (3, 'Livrée');
    INSERT INTO commande_status (status_id, description) VALUES (4, 'Annulée');

    INSERT INTO payment (paiement_id, type) VALUES (1, 'Carte bancaire');
    INSERT INTO payment (paiement_id, type) VALUES (2, 'Paypal');
    
    `
	_, err := Db.Exec(query)
	if err != nil {
		log.Fatalf("Erreur lors de l'insertion des données: %v", err)
	}

	log.Println("Données insérées avec succès")
}

func AUTO_INCREMENT() int {
	id := 0
	query := `SELECT MAX(client_id) FROM user;`
	row := Db.QueryRow(query)
	err := row.Scan(&id)
	if err != nil {
		log.Fatalf("Erreur lors de la récupération de l'ID: %v", err)
	}
	return id
}

func LoadData() {
	password := HashPassword("password")
	date := "2021-06-01 00:00:00"

	query := `
    INSERT INTO user (client_id, nom, prenom, email, telephone, hashed_password, role_id, photo_id) VALUES 
    (1, 'Mascaro', 'Giovanni', 'giovannimascaro@gmail.com', '0606060606', ?, 2, 1),
    (2, 'Elmir', 'Elias', 'eliaselmir@gmail.com', '0606060605', ?, 2, 2),
    (3, 'Zekri', 'Ilyes', 'ilyeszekri@gmail.com', '0606060603', ?, 2, 3),
    (4, 'Gari', 'Julia', 'juliagari@gmail.com', '0606060602', ?, 2, 4),
    (5, 'Ammari', 'Youssef', 'ammariyoussef@gmail.com', '0606060601', ?, 2, 5),
    (0, 'Admin', 'Admin', 'admin@gmail.com', '0606060605', ?, 1, 0),
    (6, 'Vendeur', 'Vendeur', 'vendeur@gmail.com', '0606060605', ?, 3, 6);

    INSERT INTO vendeur (vendeur_id, client_id, date_creation, nom_commercial, type_activite, n_registre_commerce, n_tva, telephone, email) VALUES
    (1, 6, ?, 'Lacoste', 'Vente de vêtements', '123456789', 'FR987654321', '0606060606', 'emailvente@gmail.com');

    INSERT INTO address (adresse_id, client_id, numero_rue, rue, ville, code_postal, pays, region, complement_adresse) VALUES
    (1, 1, 1, 'Rue de la Paix', 'Paris', '75001', 'France', 'Ile-de-France', 'Appartement 1'),
    (2, 2, 2, 'Rue de la Liberté', 'Marseille', '13001', 'France', 'Provence-Alpes-Côte d''Azur', 'Appartement 2'),
    (3, 3, 3, 'Rue de la République', 'Lyon', '69001', 'France', 'Auvergne-Rhône-Alpes', 'Appartement 3'),
    (4, 4, 4, 'Rue de la Fraternité', 'Toulouse', '31001', 'France', 'Occitanie', 'Appartement 4'),
    (5, 5, 5, 'Rue de l''Egalité', 'Nice', '06001', 'France', 'Provence-Alpes-Côte d''Azur', 'Appartement 5'),
    (6, 6, 6, 'Rue de la Justice', 'Paris', '75001', 'France', 'Ile-de-France', 'Appartement 6');

    INSERT INTO product (produit_id, nom, prix, description, stock, vendeur_id, etat_id, photo_id) VALUES
    (1, 'Polo', 50.00, 'Polo Lacoste', 100, 1, 1, 7),
    (2, 'T-shirt', 30.00, 'T-shirt Lacoste', 100, 1, 1, 8),
    (3, 'Pantalon', 70.00, 'Pantalon Lacoste', 100, 1, 1, 9),
    (4, 'Chaussures', 100.00, 'Chaussures Lacoste', 100, 1, 1, 10),
    (5, 'Casquette', 20.00, 'Casquette Lacoste', 100, 1, 1, 11),
    (6, 'Short', 40.00, 'Short Lacoste', 100, 1, 1, 12);

    INSERT INTO cart (panier_id, client_id, montant, date_creation, status_id) VALUES
    (1, 1, 0.00, ?, 1),
    (2, 2, 0.00, ?, 1),
    (3, 3, 0.00, ?, 1),
    (4, 4, 0.00, ?, 1),
    (5, 5, 0.00, ?, 1);

    INSERT INTO panier_produit (panier_id, produit_id, quantite) VALUES
    (1, 1, 1),
    (1, 2, 1),
    (1, 3, 1),
    (1, 4, 1),
    (1, 5, 1),
    (1, 6, 1),
    (2, 1, 1),
    (3, 3, 3),
    (4, 4, 4),
    (5, 5, 5);

    INSERT INTO command (commande_id, client_id, montant, status_id, date_livraison, paiement_id) VALUES
    (1, 1, 0.00, 1, ?, 1),
    (2, 2, 0.00, 1, ?, 1),
    (3, 3, 0.00, 1, ?, 1),
    (4, 4, 0.00, 1, ?, 1),
    (5, 5, 0.00, 1, ?, 1);

    INSERT INTO commande_produit (commande_id, produit_id, quantite) VALUES
    (1, 1, 1),
    (1, 2, 1),
    (1, 3, 1),
    (1, 4, 1),
    (1, 5, 1),
    (1, 6, 1),
    (2, 1, 1);

    INSERT INTO rate (evaluation_id, produit_id, client_id, note, commentaire) VALUES
    (1, 1, 1, 5, 'Très bon produit'),
    (2, 2, 1, 4, 'Bon produit'),
    (3, 3, 1, 3, 'Produit correct'),
    (4, 4, 1, 2, 'Mauvais produit'),
    (5, 5, 1, 1, 'Très mauvais produit'),
    (6, 6, 1, 5, 'Très bon produit'),
    (7, 1, 2, 4, 'Bon produit'),
    (8, 2, 2, 3, 'Produit correct'),
    (9, 3, 2, 2, 'Mauvais produit'),
    (10, 4, 2, 1, 'Très mauvais produit'),
    (11, 5, 2, 5, 'Très bon produit'),
    (12, 6, 2, 4, 'Bon produit'),
    (13, 1, 3, 3, 'Produit correct'),
    (14, 2, 3, 2, 'Mauvais produit'),
    (15, 3, 3, 1, 'Très mauvais produit'),
    (16, 4, 3, 5, 'Très bon produit'),
    (17, 5, 3, 4, 'Bon produit'),
    (18, 6, 3, 3, 'Produit correct'),
    (19, 1, 4, 2, 'Mauvais produit'),
    (20, 2, 4, 1, 'Très mauvais produit'),
    (21, 3, 4, 5, 'Très bon produit'),
    (22, 4, 4, 4, 'Bon produit'),
    (23, 5, 4, 3, 'Produit correct'),
    (24, 6, 4, 2, 'Mauvais produit'),
    (25, 1, 5, 1, 'Très mauvais produit'),
    (26, 2, 5, 5, 'Très bon produit'),
    (27, 3, 5, 4, 'Bon produit'),
    (28, 4, 5, 3, 'Produit correct'),
    (29, 5, 5, 2, 'Mauvais produit'),
    (30, 6, 5, 1, 'Très mauvais produit');

    INSERT INTO photo (photo_id, url) VALUES
    (1, 'https://www.lacoste.com/dw/image/v2/AAQM_PRD/on/demandware.static/-/Sites/default/dw7b3b3b3b/images/hi-res/EF8470_51_001.jpg'),
    (2, 'https://www.lacoste.com/dw/image/v2/AAQM_PRD/on/demandware.static/-/Sites/default/dw7b3b3b3b/images/hi-res/EF8470_51_001.jpg'),
    (3, 'https://www.lacoste.com/dw/image/v2/AAQM_PRD/on/demandware.static/-/Sites/default/dw7b3b3b3b/images/hi-res/EF8470_51_001.jpg'),
    (4, 'https://www.lacoste.com/dw/image/v2/AAQM_PRD/on/demandware.static/-/Sites/default/dw7b3b3b3b/images/hi-res/EF8470_51_001.jpg'),
    (5, 'https://www.lacoste.com/dw/image/v2/AAQM_PRD/on/demandware.static/-/Sites/default/dw7b3b3b3b/images/hi-res/EF8470_51_001.jpg');
    `

	_, err := Db.Exec(query,
		password, password, password, password, password, password, password,
		date, date, date, date, date, date, date, date, date, date, date, date,
	)
	if err != nil {
		log.Fatalf("Erreur lors de l'insertion des données: %v", err)
	}

	log.Println("Données insérées avec succès")
}
