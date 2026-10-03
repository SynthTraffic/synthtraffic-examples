# Healthcare fixtures

Fictional data for testing healthcare pipelines. No patient source files.
Observations, services, prices, and claim outcomes are illustrative fixture
parameters, not clinical models. Custom JSON records — not FHIR or HL7.

| Scenario | Behavior |
| --- | --- |
| [`patient_encounters.yaml`](patient_encounters.yaml) | Patients, practitioners, encounter lifecycle with refs |
| [`bedside_observations.yaml`](bedside_observations.yaml) | Device instances with `previous` heart rate |
| [`claims.yaml`](claims.yaml) | Stable line items, then weighted paid/denied |

Each scenario pins seed `42` and a UTC clock start.

```bash
synthtraffic sample healthcare/patient_encounters.yaml --events 19 --seed 42
synthtraffic run healthcare/bedside_observations.yaml --stdout --seed 42
synthtraffic run healthcare/claims.yaml --stdout --seed 42
```

Optional pre-generated JSONL lives under [`samples/`](samples/) when present.
