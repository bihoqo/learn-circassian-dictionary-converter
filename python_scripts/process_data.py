import os
import re

# =============================================================================
# CONFIGURATION
# =============================================================================
# 1. Determine the directory where this script is located (/python_scripts/)
SCRIPT_DIR = os.path.dirname(os.path.abspath(__file__))

# 2. Define Relative Paths
# Input: backup folder
BACKUP_DIR = os.path.join(SCRIPT_DIR, '..', 'content', 'backup')
# Output: phase-01-raw-data folder
OUTPUT_DIR = os.path.join(SCRIPT_DIR, '..', 'content', 'phase-01-raw-data')

# File Definitions
FILE_1_NAME = '30-Ady-Rus_ThreeVolumes.txt'
FILE_2_NAME = '33-Ady-Rus-1960.txt'

# =============================================================================
# LOGIC: FILE 1 (30-Ady-Rus_ThreeVolumes.txt)
# =============================================================================
def process_file_1(input_path, output_path):
    print(f"--- Processing File 1: {os.path.basename(input_path)} ---")
    change_log = []

    # Logic:
    # 1. Add space before dot suffix: "word1." -> "word 1."
    # 2. Add space before comma suffix: "word2," -> "word 2,"
    # 3. Remove number before slash: "word1/" -> "word /"
    # Lookbehind ensures we don't match digits preceded by space or '('.
    pattern = re.compile(r'(?<=[^\s\d\(])(\d+\.|\d+,|\d+\s*\/)')

    def callback(match, line_num):
        original = match.group(0)

        if original.endswith('.'):
            # Case: "1." -> " 1."
            new_text = " " + original
            action = "Added Space (Dot)"
        elif original.endswith(','):
            # Case: "2," -> " 2,"
            new_text = " " + original
            action = "Added Space (Comma)"
        elif '/' in original:
            # Case: "1/" -> " /"
            new_text = " /"
            action = "Removed Number (Slash)"
        else:
            return original

        change_log.append(f"Line {line_num}: '{original}' -> '{new_text}' ({action})")
        return new_text

    try:
        os.makedirs(os.path.dirname(output_path), exist_ok=True)

        with open(input_path, 'r', encoding='utf-8') as f_in, \
             open(output_path, 'w', encoding='utf-8') as f_out:

            for line_num, line in enumerate(f_in, 1):
                new_line = pattern.sub(lambda m: callback(m, line_num), line)
                f_out.write(new_line)

        print(f"Saved to: {output_path}")
        print(f"Total Changes: {len(change_log)}")

    except FileNotFoundError:
        print(f"Error: Could not find {input_path}")
    except Exception as e:
        print(f"Error processing File 1: {e}")

# =============================================================================
# LOGIC: FILE 2 (33-Ady-Rus-1960.txt)
# =============================================================================
def process_file_2(input_path, output_path):
    print(f"\n--- Processing File 2: {os.path.basename(input_path)} ---")
    change_log = []

    # Regex to find the start of the definition (first lowercase Cyrillic or paren)
    split_pattern = re.compile(r'[\(\[а-я]')

    try:
        os.makedirs(os.path.dirname(output_path), exist_ok=True)

        with open(input_path, 'r', encoding='utf-8') as f_in, \
             open(output_path, 'w', encoding='utf-8') as f_out:

            for line_num, line in enumerate(f_in, 1):
                # 1. Identify the Headword Block
                match = split_pattern.search(line)

                if match:
                    split_index = match.start()
                    head_part = line[:split_index]
                    rest_part = line[split_index:]

                    # Only process if head_part isn't empty and likely contains a headword
                    if head_part.strip():
                        original_head = head_part

                        # STEP A: Remove ALL spaces (Fixes "И Т У М" -> "ИТУМ")
                        # This also turns "WORD 1." into "WORD1." (temporarily)
                        cleaned_head = head_part.replace(" ", "")

                        # STEP B: Re-insert space before number suffix
                        # Regex explanation:
                        # (\d+)    : Capture the number at the end
                        # ([\.,]?) : Capture optional dot or comma following it
                        # $        : End of string
                        if re.search(r'\d[\.,]?$', cleaned_head):
                             cleaned_head = re.sub(r'(\d+)([\.,]?)$', r' \1\2', cleaned_head)

                        # Reconstruct the line
                        # We use .lstrip() on rest_part to ensure clean separation
                        if cleaned_head != original_head:
                            new_line = f"{cleaned_head} {rest_part.lstrip()}"

                            # Only log if it's a substantive change (ignoring purely whitespace alignment)
                            if cleaned_head.strip() != original_head.strip().replace("  ", " "):
                                change_log.append(f"Line {line_num}: '{original_head.strip()}' -> '{cleaned_head}'")

                            f_out.write(new_line)
                            continue

                # Fallback: Write original line if no changes needed
                f_out.write(line)

        print(f"Saved to: {output_path}")
        print(f"Total Changes: {len(change_log)}")
        if change_log:
            print("Sample Changes:")
            for log in change_log[:5]:
                print(f"  {log}")

    except FileNotFoundError:
        print(f"Error: Could not find {input_path}")
    except Exception as e:
        print(f"Error processing File 2: {e}")

# =============================================================================
# MAIN EXECUTION
# =============================================================================
if __name__ == "__main__":
    print(f"Running from: {os.getcwd()}")

    # 1. Process 30-Ady-Rus_ThreeVolumes.txt
    f1_in = os.path.join(BACKUP_DIR, FILE_1_NAME)
    f1_out = os.path.join(OUTPUT_DIR, FILE_1_NAME)
    process_file_1(f1_in, f1_out)

    # 2. Process 33-Ady-Rus-1960.txt
    f2_in = os.path.join(BACKUP_DIR, FILE_2_NAME)
    f2_out = os.path.join(OUTPUT_DIR, FILE_2_NAME)
    process_file_2(f2_in, f2_out)