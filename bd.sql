CREATE DATABASE IF NOT EXISTS Veterinary;
USE Veterinary;

CREATE USER IF NOT EXISTS 'rubi'@'localhost' IDENTIFIED BY '1234';
GRANT ALL PRIVILEGES ON Veterinary.* TO 'rubi'@'localhost';
FLUSH PRIVILEGES;

CREATE TABLE IF NOT EXISTS user (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS pet (
    id INT AUTO_INCREMENT PRIMARY KEY,
    `group` VARCHAR(100),
    control VARCHAR(100),
    race VARCHAR(100),
    age INT,
    gender VARCHAR(20),
    weight VARCHAR(20),
    bodyCondition INT,
    diagnosis TEXT,
    degreeLameness INT,
    onsetTimeSymptoms DATE,
    name VARCHAR(100),
    owner VARCHAR(150),
    color VARCHAR(100),
    lastAppointment DATE,
    image VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS appointment (
    id INT AUTO_INCREMENT PRIMARY KEY,
    pet_id INT,
    date DATETIME,
    FOREIGN KEY (pet_id) REFERENCES pet(id)
);

CREATE TABLE IF NOT EXISTS dataSense (
    id INT AUTO_INCREMENT PRIMARY KEY,
    idAppointment INT,
    type VARCHAR(20),
    totalTime VARCHAR(50),
    frequencyHZ INT,
    amplitudeMV INT,
    COPN TEXT,
    COPC TEXT,
    result TEXT,
    gyroscope JSON,
    accelerometer JSON,
    symmetryIndexLF TEXT,
    symmetryIndexRF TEXT,
    symmetryIndexLB TEXT,
    symmetryIndexRB TEXT,
    weightDistributionLF TEXT,
    weightDistributionRF TEXT,
    weightDistributionLB TEXT,
    weightDistributionRB TEXT,
    verticalForce TEXT,
    verticalImpulse LONGTEXT,
    FOREIGN KEY (idAppointment) REFERENCES appointment(id)
);

INSERT INTO user (name, password) VALUES
('Leydi Rubí', '1234'),
('Isis Naomi', '1234');

INSERT INTO pet (`group`, control, race, age, gender, weight, bodyCondition, 
                diagnosis, degreeLameness, onsetTimeSymptoms, name, owner, 
                color, lastAppointment, image) VALUES
('Prueba', 'Hola mundo', 'Bull Terrier', 7, 'Macho', '35', 3, 
 'Prueba (Extremidades afectadas: PI)', 1, '2026-02-19', 'Bull', 'Adrián', 
 'Café', '2026-02-27', 'https://www.purina.es/sites/default/files/styles/ttt_image_510/public/2021-01/Bull%20Terrier1.jpg?itok=kdIZROtx');

INSERT INTO appointment (pet_id, date) VALUES
(3, '2026-02-22 06:36:00');

INSERT INTO dataSense (idAppointment, type, totalTime, frequencyHZ, amplitudeMV, 
                      COPN, COPC, result, gyroscope, accelerometer,
                      symmetryIndexLF, symmetryIndexRF, symmetryIndexLB, symmetryIndexRB,
                      weightDistributionLF, weightDistributionRF, weightDistributionLB, weightDistributionRB,
                      verticalForce, verticalImpulse) VALUES
(1, 'Análisis marcha', '30 seg', 100, 5,
 '{"x": 2.5, "y": 1.8}', '{"x": 2.3, "y": 1.9}', 'Normal',
 '{"x": 0.5, "y": 0.3, "z": 0.2}', '{"x": 9.8, "y": 0.1, "z": 0.1}',
 '0.95', '0.96', '0.94', '0.95',
 '28%', '27%', '23%', '22%',
 '450 N', '12.5 N·s');

SELECT 'Tabla user:' as '';
SELECT * FROM user;

SELECT 'Tabla pet:' as '';
SELECT * FROM pet;

SELECT 'Tabla appointment:' as '';
SELECT a.*, p.name as pet_name, p.owner 
FROM appointment a 
LEFT JOIN pet p ON a.pet_id = p.id;

SELECT 'Tabla dataSense:' as '';
SELECT ds.*, a.date as appointment_date, p.name as pet_name
FROM dataSense ds 
LEFT JOIN appointment a ON ds.idAppointment = a.id
LEFT JOIN pet p ON a.pet_id = p.id;