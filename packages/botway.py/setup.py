from setuptools import setup
import re

version = ""
with open("botway/__init__.py") as f:
    version = re.search(r'^__version__\s*=\s*[\'"]([^\'"]*)[\'"]', f.read(), re.MULTILINE).group(1)

if not version:
    raise RuntimeError("version is not set")

readme = ""
with open("README.md") as f:
    readme = f.read()

setup(
    name="botway.py",
    author="abdfnx",
    url="https://github.com/abdfnx/botway",
    project_urls={
        "Issues": "https://github.com/abdfnx/botway/issues",
    },
    version=version,
    packages=["botway"],
    license="MIT",
    description="Python client package for Botway.",
    long_description=readme,
    long_description_content_type="text/markdown",
    include_package_data=True,
    python_requires=">=3.6.0",
    classifiers=[
        "Development Status :: 5 - Production/Stable",
        "License :: OSI Approved :: MIT License",
        "Intended Audience :: Developers",
        "Natural Language :: English",
        "Operating System :: OS Independent",
        "Programming Language :: Python :: 3.8",
        "Programming Language :: Python :: 3.9",
        "Programming Language :: Python :: 3.10",
        "Topic :: Internet",
        "Topic :: Software Development :: Libraries",
        "Topic :: Software Development :: Libraries :: Python Modules",
        "Topic :: Utilities",
    ]
)
