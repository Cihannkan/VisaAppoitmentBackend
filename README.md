# VisaAppoitmentBackend

VisaAppoitmentBackend, Schengen vizesi randevularını yönetmek için bir backend uygulamasıdır. Bu uygulama, belirli bir ülkeye göre vize randevularını filtrelemenize ve mevcut randevu bilgilerini almanıza olanak tanır.

## Özellikler

- Mevcut vize randevularını listeleme
- Belirli bir ülkeye göre randevuları filtreleme

## Kurulum

1. Bu projeyi klonlayın:
   ```bash
   git clone https://github.com/Cihannkan/VisaAppoitmentBackend.git
   cd VisaAppoitmentBackend
   ```

2. Gerekli bağımlılıkları yükleyin:
   ```bash
   go mod tidy
   ```

3. Uygulamayı başlatın:
   ```bash
   go run main.go
   ```

## Kullanım

- Tüm randevuları listelemek için:
  ```
  GET /
  ```
  
- Tüm misyon ülkelerini listelemek için:
  ```
  GET /getmissioncountrylist
  ```

- Belirli bir ülkeye göre randevuları almak için:
  ```
  GET /getappointmentbycountry/:country
  ```
