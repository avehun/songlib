CREATE TABLE IF NOT EXISTS songs (
    id SERIAL PRIMARY KEY,
    "group" VARCHAR(255) NOT NULL,
    song VARCHAR(255) NOT NULL,
    release_date TEXT NOT NULL,
    "text" TEXT NOT NULL,
    link VARCHAR(255) NOT NULL
);

-- Nota bene: I generated examples with ChatGPT
INSERT INTO songs ("group", song, release_date, "text", link)
VALUES 
    ('The Beatles', 'Hey Jude', '1968-08-26', 'Hey Jude, dont make it bad. Take a sad song and make it better.', 'https://example.com/hey-jude'),
    ('Queen', 'Bohemian Rhapsody', '1975-10-31', 'Is this the real life? Is this just fantasy?', 'https://example.com/bohemian-rhapsody'),
    ('Pink Floyd', 'Comfortably Numb', '1979-11-30', 'Hello? Is there anybody in there? Just nod if you can hear me.', 'https://example.com/comfortably-numb'),
    ('Adele', 'Rolling in the Deep', '2010-11-29', 'We could have had it all.', 'https://example.com/rolling-in-the-deep'),
    ('Led Zeppelin', 'Stairway to Heaven', '1971-11-08', 'Theres a lady whos sure all that glitters is gold', 'https://example.com/stairway-to-heaven');