import re

# Nama file SQL yang ingin diubah
input_file = "indonesia_sub_districts.sql"      # Ganti dengan nama file asli
output_file = "ind_sub_district.sql"  # File hasil perubahan

# Baca isi file
with open(input_file, "r", encoding="utf-8") as file:
    sql_content = file.read()

# Ganti semua "indonesia_cities" menjadi "indonesia_districts"
updated_sql = re.sub(r"\bindonesia_districts\b", "indonesia_sub_districts", sql_content)

# Simpan hasil perubahan ke file baru
with open(output_file, "w", encoding="utf-8") as file:
    file.write(updated_sql)

print(f"Perubahan selesai! File hasil: {output_file}")
